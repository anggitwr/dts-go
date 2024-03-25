package router

import (
	"finalpro/controller"
	"finalpro/middleware"
	"finalpro/repository"
	"finalpro/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartRouter(r *gin.Engine, db *gorm.DB) {

	repoUser := repository.NewUserRepository(db)
	srcUser := service.NewUserService(repoUser)
	ctrlUser := controller.NewUserController(srcUser)
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", ctrlUser.Register)

		userRouter.POST("/login", ctrlUser.Login)
		userRouter.PUT("/:userId", middleware.Authentication(), ctrlUser.UpdateUser)
		userRouter.DELETE("/:userId", middleware.Authentication(), ctrlUser.DeleteUser)
	}

	repoPhoto := repository.NewPhotoRepository(db)
	srcPhoto := service.NewPhotoService(repoPhoto)
	ctrlPhoto := controller.NewPhotoController(srcPhoto)
	photoRouter := r.Group("/photo")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.GET("/", ctrlPhoto.GetAllPhotos)
		photoRouter.GET("/:photoId", ctrlPhoto.GetPhotoByID)
		photoRouter.POST("/", ctrlPhoto.CreatePhoto)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), ctrlPhoto.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), ctrlPhoto.DeletePhoto)
	}
	repoComment := repository.NewCommentRepository(db)
	srcComment := service.NewCommentService(repoComment)
	ctrlComment := controller.NewCommentController(srcComment)
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.GET("/", ctrlComment.GetAllComments)
		commentRouter.GET("/:commentId", ctrlComment.GetCommentByID)
		commentRouter.POST("/", ctrlComment.CreateComment)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), ctrlComment.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), ctrlComment.DeleteComment)
	}
	repoSosmed := repository.NewSosmedRepository(db)
	srcSosmed := service.NewSosmedService(repoSosmed)
	ctrlSosmed := controller.NewSosmedController(srcSosmed)
	// Social Media
	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Authentication())
		socialMediaRouter.GET("/", ctrlSosmed.GetAllSosmed)
		socialMediaRouter.GET("/:socialMediaId", ctrlSosmed.GetSosmedByID)
		socialMediaRouter.POST("/", ctrlSosmed.CreateSosmed)
		socialMediaRouter.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), ctrlSosmed.UpdateSosmed)
		socialMediaRouter.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), ctrlSosmed.DeleteSosmed)
	}
}
