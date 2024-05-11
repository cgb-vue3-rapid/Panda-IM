package Auth

import (
	"akita/panda-im/common/constants"
	"akita/panda-im/common/models/ctype"
	"akita/panda-im/service/user/api/code"
	"akita/panda-im/service/user/api/internal/svc"
	"akita/panda-im/service/user/api/internal/types"
	"akita/panda-im/service/user/rpc/models/entity"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
)

type UserConfigUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserConfigUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserConfigUpdateLogic {
	return &UserConfigUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserConfigUpdateLogic) UserConfigUpdate(req *types.UserConfigUpdateRequest) (resp *types.UserConfigUpdateResponse, err error) {
	// todo: add your logic here and delete this line
	//获取用户ID
	userId := l.ctx.Value(constants.UserId).(int64)

	// 获取用户信息
	userMap := RefToMap(*req, "user")
	fmt.Println(userMap)
	if len(userMap) != 0 {
		var userModel entity.UserModel
		err := l.svcCtx.Orm.WithContext(l.ctx).Model(&userModel).First(&userModel, userId).Error
		if err != nil {
			logx.Errorf("Failed to get user: %v", err)
			return nil, code.ErrUpdateFailed
		}

		err = l.svcCtx.Orm.Model(&userModel).Updates(userMap).Error
		if err != nil {
			return nil, code.ErrUpdateFailed
		}
	}

	// 获取用户配置信息
	userConfigMap := RefToMap(*req, "user_config")

	if len(userConfigMap) != 0 {
		var userConfigModel entity.UserConfModel
		err := l.svcCtx.Orm.WithContext(l.ctx).Model(&userConfigModel).First(&userConfigModel, userId).Error
		if err != nil {
			return nil, code.ErrUpdateFailed
		}

		verificationQuestion, ok := userConfigMap["verification_question"]
		if ok {
			delete(userConfigMap, "verification_question")
			data := ctype.VerificationQuestion{}
			MapToStruct(verificationQuestion.(map[string]any), &data)
			l.svcCtx.Orm.Model(&userConfigModel).Updates(&entity.UserConfModel{VerificationQuestion: &data})
		}
		l.svcCtx.Orm.Model(&userConfigModel).Updates(userConfigMap)
	}

	return &types.UserConfigUpdateResponse{Message: "更新成功"}, nil
}

func RefToMap(data interface{}, tag string) map[string]interface{} {
	maps := make(map[string]interface{})
	v := reflect.ValueOf(data)

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		getTag, ok := field.Tag.Lookup(tag)
		if !ok {
			continue
		}
		val := v.Field(i)
		if val.IsZero() {
			continue
		}

		if field.Type.Kind() == reflect.Struct {
			newMaps := RefToMap(val.Interface(), tag)
			maps[getTag] = newMaps
			continue
		}

		if field.Type.Kind() == reflect.Ptr {
			if field.Type.Elem().Kind() == reflect.Struct {
				newMaps := RefToMap(val.Elem().Interface(), tag)
				maps[getTag] = newMaps
				continue
			}
			maps[getTag] = val.Elem().Interface()
			continue
		}

		maps[getTag] = val.Interface()
	}
	return maps
}

func MapToStruct(data map[string]interface{}, dst interface{}) {
	// 确保dst是一个指针
	dstPtr := reflect.ValueOf(dst)
	if dstPtr.Kind() != reflect.Ptr {
		panic("dst must be a pointer")
	}

	// 获取指针所指向的结构体类型和值
	t := reflect.TypeOf(dst).Elem()
	v := reflect.ValueOf(dst).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue // 跳过没有JSON标签或标签为"-"的字段
		}

		mapField, ok := data[tag]
		if !ok {
			continue // 跳过在数据中找不到的字段
		}

		// 获取字段值
		val := v.Field(i)

		// 如果字段类型是指针
		if field.Type.Kind() == reflect.Ptr {
			// 获取指针指向的类型
			ptrType := field.Type.Elem()

			// 如果映射的值是字符串类型
			if ptrType.Kind() == reflect.String {
				if mapFieldValue, ok := mapField.(string); ok {
					// 创建新的字符串值并设置字段的指针
					newVal := reflect.New(ptrType)
					newVal.Elem().SetString(mapFieldValue)
					val.Set(newVal)
				}
			}
			// 根据需要添加更多类型的情况
		}
	}
}
