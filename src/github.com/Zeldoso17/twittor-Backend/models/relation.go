package models

type Relation struct {
	UserID 			  string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuariorelacionId"`
}