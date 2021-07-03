package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"

	"github.com/childelins/go-gin-api/config"
	"github.com/childelins/go-gin-api/pkg/logger"
	"github.com/childelins/go-gin-api/pkg/registry"
	"github.com/childelins/go-gin-api/proto"
)

var (
	ServerConfig *config.ServerConfig
	NacosConfig  *config.NacosConfig
	Logger       *logger.Logger
	Trans        ut.Translator
	DB           *gorm.DB
	Tracer       opentracing.Tracer
	Registry     registry.Registry

	LecturerSrvClient proto.LecturerClient
)
