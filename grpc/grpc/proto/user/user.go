package user

import context "context"

//定义服务端实现约定的接口
type UserInfoService struct {
	UnimplementedUserInfoServiceServer
}

func NewUserInfoService() *UserInfoService {
	return &UserInfoService{}
}

//实现服务端需要实现的接口
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *UserRequest) (resp *UserResponse, err error) {
	name := req.Name
	//在库中查用户信息
	if name == "zs" {
		resp = &UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Hobby: []string{"sing", "run", "eat"},
		}
	}
	return
}
