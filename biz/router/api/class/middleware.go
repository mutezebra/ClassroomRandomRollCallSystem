// Code generated by hertz generator.

package class

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/mutezebra/ClassroomRandomRollCallSystem/biz/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _authMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.JWT()}
}

func _createclassMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _joinclassMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _classMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _classlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _classstudentlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _viewinvitationcodeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getclassteacherMw() []app.HandlerFunc {
	// your code...
	return nil
}
