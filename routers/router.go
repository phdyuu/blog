package routers

import (
	"phdyu/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "get:ShowIndex")
    beego.Router("/loginyu", &controllers.LoginController{}, "get:ShowAdminIndex;post:HandleLogin")
	beego.Router("/friends", &controllers.PhotoController{}, "get:ShowFriends")
	beego.Router("/personal", &controllers.PhotoController{}, "get:ShowPersonal")
	beego.Router("/cats", &controllers.PhotoController{}, "get:ShowCats")
	beego.Router("/scenery", &controllers.PhotoController{}, "get:ShowScenery")
	beego.Router("/portfolio", &controllers.PhotoController{}, "get:ShowPhotoClass")
	beego.Router("/index", &controllers.MainController{}, "get:ShowIndex")
	beego.Router("/AddMoment", &controllers.AdminController{}, "get:ShowAddMoment;post:HandleAddMoment")
    beego.Router("/UpdateMoment", &controllers.AdminController{}, "get:ShowUpdateMoment;post:HandleUpdateMoment")
	beego.Router("/ShowAdminSystem", &controllers.AdminController{}, "get:ShowAdminSystem;post:HandleAddMoment")
	beego.Router("/ShowAdminAlbum", &controllers.AdminController{}, "get:ShowAdminAlbum")
    beego.Router("/DeleteMoment", &controllers.AdminController{}, "post:HandleDeleteMoment")


}
