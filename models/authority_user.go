package models

type SysUser struct {
	ID        uint   `json:"ID" gorm:"comment:'用户ID';AUTO_INCREMENT"`
	Username  string `json:"username" gorm:"comment:'用户登录名'not null;unique;"`
	Password  string `json:"-"  gorm:"comment:'用户登录密码'"`
	NickName  string `json:"nickName" gorm:"default:'系统用户';comment:'用户昵称'" `
	HeaderImg string `json:"headerImg" gorm:"default:'http://qmplusimg.henrongyi.top/head.png';comment:'用户头像'"`
}

func (s *SysUser) TableName() string {
	return "users"
}

func (s *SysUser) Login() (err error) {
	err = DB.Where("username = ? AND password = ?", s.Username, s.Password).First(s).Error
	return err
}

func (s *SysUser) Update(UpDataPassword string) (err error) {
	err = DB.Where("ID = ? AND password = ?", s.ID, s.Password).First(s).Error
	if err != nil {
		return err
	} else {
		s.Password = UpDataPassword
		err = DB.Save(s).Error
	}
	return err
}

func (s *SysUser) Create() (err error) {
	s.NickName = s.Username
	err = DB.Create(s).Error
	return err
}
