package server

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"

	"github.com/photoprism/photoprism/internal/auto"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/pkg/clean"
	"github.com/photoprism/photoprism/pkg/fs"
)

const WebDAVOriginals = "/originals"
const WebDAVImport = "/import"

// registerWebDAVRoutes configures the built-in WebDAV server.
func registerWebDAVRoutes(router *gin.Engine, conf *config.Config) {
	if conf.DisableWebDAV() {
		log.Info("webdav: server disabled")
	} else {
		WebDAV(conf.OriginalsPath(), router.Group(conf.BaseUri(WebDAVOriginals), BasicAuth()), conf)
		log.Infof("webdav: %s/ enabled, waiting for requests", conf.BaseUri(WebDAVOriginals))

		if conf.ImportPath() != "" {
			WebDAV(conf.ImportPath(), router.Group(conf.BaseUri(WebDAVImport), BasicAuth()), conf)
			log.Infof("webdav: %s/ enabled, waiting for requests", conf.BaseUri(WebDAVImport))
		}
	}
}

var WebDAVHandler = func(c *gin.Context, router *gin.RouterGroup, srv *webdav.Handler) {
	srv.ServeHTTP(c.Writer, c.Request)
}

// WebDAV handles requests to the /originals and /import endpoints.
func WebDAV(filePath string, router *gin.RouterGroup, conf *config.Config) {
	if router == nil {
		log.Error("webdav: router is nil")
		return
	}

	if conf == nil {
		log.Error("webdav: conf is nil")
		return
	}

	// Native file system restricted to a specific directory.
	fileSystem := webdav.Dir(filePath)

	// Request logger function.
	loggerFunc := func(r *http.Request, err error) {
		if err != nil {
			switch r.Method {
			case MethodPut, MethodPost, MethodPatch, MethodDelete, MethodCopy, MethodMove:
				log.Errorf("webdav: %s in %s %s", clean.Log(err.Error()), clean.Log(r.Method), clean.Log(r.URL.String()))
			case MethodPropfind:
				log.Tracef("webdav: %s in %s %s", clean.Log(err.Error()), clean.Log(r.Method), clean.Log(r.URL.String()))
			default:
				log.Debugf("webdav: %s in %s %s", clean.Log(err.Error()), clean.Log(r.Method), clean.Log(r.URL.String()))
			}
		} else {
			// Mark uploaded files as favorite if X-Favorite HTTP header is "1".
			if r.Method == MethodPut && r.Header.Get("X-Favorite") == "1" {
				if router.BasePath() == conf.BaseUri(WebDAVOriginals) {
					MarkUploadAsFavorite(filepath.Join(conf.OriginalsPath(), strings.TrimPrefix(r.URL.Path, router.BasePath())))
				} else if router.BasePath() == conf.BaseUri(WebDAVImport) {
					MarkUploadAsFavorite(filepath.Join(conf.ImportPath(), strings.TrimPrefix(r.URL.Path, router.BasePath())))
				}
			}

			switch r.Method {
			case MethodPut, MethodPost, MethodPatch, MethodDelete, MethodCopy, MethodMove:
				log.Infof("webdav: %s %s", clean.Log(r.Method), clean.Log(r.URL.String()))

				if router.BasePath() == conf.BaseUri(WebDAVOriginals) {
					auto.ShouldIndex()
				} else if router.BasePath() == conf.BaseUri(WebDAVImport) {
					auto.ShouldImport()
				}
			default:
				log.Tracef("webdav: %s %s", clean.Log(r.Method), clean.Log(r.URL.String()))
			}
		}
	}

	// WebDAV request handler.
	srv := &webdav.Handler{
		Prefix:     router.BasePath(),
		FileSystem: fileSystem,
		LockSystem: webdav.NewMemLS(),
		Logger:     loggerFunc,
	}

	// Request handler wrapper function.
	handlerFunc := func(c *gin.Context) {
		WebDAVHandler(c, router, srv)
	}

	// Handle supported HTTP request methods.
	router.Handle(MethodHead, "/*path", handlerFunc)
	router.Handle(MethodGet, "/*path", handlerFunc)
	router.Handle(MethodPut, "/*path", handlerFunc)
	router.Handle(MethodPost, "/*path", handlerFunc)
	router.Handle(MethodPatch, "/*path", handlerFunc)
	router.Handle(MethodDelete, "/*path", handlerFunc)
	router.Handle(MethodOptions, "/*path", handlerFunc)
	router.Handle(MethodMkcol, "/*path", handlerFunc)
	router.Handle(MethodCopy, "/*path", handlerFunc)
	router.Handle(MethodMove, "/*path", handlerFunc)
	router.Handle(MethodLock, "/*path", handlerFunc)
	router.Handle(MethodUnlock, "/*path", handlerFunc)
	router.Handle(MethodPropfind, "/*path", handlerFunc)
	router.Handle(MethodProppatch, "/*path", handlerFunc)
}

// MarkUploadAsFavorite sets the favorite flag for newly uploaded files.
func MarkUploadAsFavorite(fileName string) {
	yamlName := fs.AbsPrefix(fileName, false) + fs.ExtYAML

	// Abort if YAML file already exists to avoid overwriting metadata.
	if fs.FileExists(yamlName) {
		log.Warnf("webdav: %s already exists", clean.Log(filepath.Base(yamlName)))
		return
	}

	// Make sure directory exists.
	if err := os.MkdirAll(filepath.Dir(yamlName), os.ModePerm); err != nil {
		log.Errorf("webdav: %s", err.Error())
		return
	}

	// Write YAML data to file.
	if err := os.WriteFile(yamlName, []byte("Favorite: true\n"), os.ModePerm); err != nil {
		log.Errorf("webdav: %s", err.Error())
		return
	}

	// Log success.
	log.Infof("webdav: marked %s as favorite", clean.Log(filepath.Base(fileName)))
}
