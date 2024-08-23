// Code generated by ent, DO NOT EDIT.

package avatar

import (
	"time"
	"toktok-backend/internal/adapter/persistence/database/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Avatar {
	return predicate.Avatar(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Avatar {
	return predicate.Avatar(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Avatar {
	return predicate.Avatar(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Avatar {
	return predicate.Avatar(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Avatar {
	return predicate.Avatar(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Avatar {
	return predicate.Avatar(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Avatar {
	return predicate.Avatar(sql.FieldLTE(FieldID, id))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldDeletedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldUpdatedAt, v))
}

// Birthday applies equality check predicate on the "birthday" field. It's identical to BirthdayEQ.
func Birthday(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldBirthday, v))
}

// Mbti applies equality check predicate on the "mbti" field. It's identical to MbtiEQ.
func Mbti(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldMbti, v))
}

// Picture applies equality check predicate on the "picture" field. It's identical to PictureEQ.
func Picture(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldPicture, v))
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldNickname, v))
}

// Introduce applies equality check predicate on the "introduce" field. It's identical to IntroduceEQ.
func Introduce(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldIntroduce, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Avatar {
	return predicate.Avatar(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Avatar {
	return predicate.Avatar(sql.FieldNotNull(FieldDeletedAt))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Avatar {
	return predicate.Avatar(sql.FieldLTE(FieldUpdatedAt, v))
}

// SexEQ applies the EQ predicate on the "sex" field.
func SexEQ(v Sex) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldSex, v))
}

// SexNEQ applies the NEQ predicate on the "sex" field.
func SexNEQ(v Sex) predicate.Avatar {
	return predicate.Avatar(sql.FieldNEQ(FieldSex, v))
}

// SexIn applies the In predicate on the "sex" field.
func SexIn(vs ...Sex) predicate.Avatar {
	return predicate.Avatar(sql.FieldIn(FieldSex, vs...))
}

// SexNotIn applies the NotIn predicate on the "sex" field.
func SexNotIn(vs ...Sex) predicate.Avatar {
	return predicate.Avatar(sql.FieldNotIn(FieldSex, vs...))
}

// BirthdayEQ applies the EQ predicate on the "birthday" field.
func BirthdayEQ(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldBirthday, v))
}

// BirthdayNEQ applies the NEQ predicate on the "birthday" field.
func BirthdayNEQ(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldNEQ(FieldBirthday, v))
}

// BirthdayIn applies the In predicate on the "birthday" field.
func BirthdayIn(vs ...string) predicate.Avatar {
	return predicate.Avatar(sql.FieldIn(FieldBirthday, vs...))
}

// BirthdayNotIn applies the NotIn predicate on the "birthday" field.
func BirthdayNotIn(vs ...string) predicate.Avatar {
	return predicate.Avatar(sql.FieldNotIn(FieldBirthday, vs...))
}

// BirthdayGT applies the GT predicate on the "birthday" field.
func BirthdayGT(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldGT(FieldBirthday, v))
}

// BirthdayGTE applies the GTE predicate on the "birthday" field.
func BirthdayGTE(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldGTE(FieldBirthday, v))
}

// BirthdayLT applies the LT predicate on the "birthday" field.
func BirthdayLT(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldLT(FieldBirthday, v))
}

// BirthdayLTE applies the LTE predicate on the "birthday" field.
func BirthdayLTE(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldLTE(FieldBirthday, v))
}

// BirthdayContains applies the Contains predicate on the "birthday" field.
func BirthdayContains(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldContains(FieldBirthday, v))
}

// BirthdayHasPrefix applies the HasPrefix predicate on the "birthday" field.
func BirthdayHasPrefix(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldHasPrefix(FieldBirthday, v))
}

// BirthdayHasSuffix applies the HasSuffix predicate on the "birthday" field.
func BirthdayHasSuffix(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldHasSuffix(FieldBirthday, v))
}

// BirthdayEqualFold applies the EqualFold predicate on the "birthday" field.
func BirthdayEqualFold(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEqualFold(FieldBirthday, v))
}

// BirthdayContainsFold applies the ContainsFold predicate on the "birthday" field.
func BirthdayContainsFold(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldContainsFold(FieldBirthday, v))
}

// MbtiEQ applies the EQ predicate on the "mbti" field.
func MbtiEQ(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldMbti, v))
}

// MbtiNEQ applies the NEQ predicate on the "mbti" field.
func MbtiNEQ(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldNEQ(FieldMbti, v))
}

// MbtiIn applies the In predicate on the "mbti" field.
func MbtiIn(vs ...string) predicate.Avatar {
	return predicate.Avatar(sql.FieldIn(FieldMbti, vs...))
}

// MbtiNotIn applies the NotIn predicate on the "mbti" field.
func MbtiNotIn(vs ...string) predicate.Avatar {
	return predicate.Avatar(sql.FieldNotIn(FieldMbti, vs...))
}

// MbtiGT applies the GT predicate on the "mbti" field.
func MbtiGT(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldGT(FieldMbti, v))
}

// MbtiGTE applies the GTE predicate on the "mbti" field.
func MbtiGTE(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldGTE(FieldMbti, v))
}

// MbtiLT applies the LT predicate on the "mbti" field.
func MbtiLT(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldLT(FieldMbti, v))
}

// MbtiLTE applies the LTE predicate on the "mbti" field.
func MbtiLTE(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldLTE(FieldMbti, v))
}

// MbtiContains applies the Contains predicate on the "mbti" field.
func MbtiContains(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldContains(FieldMbti, v))
}

// MbtiHasPrefix applies the HasPrefix predicate on the "mbti" field.
func MbtiHasPrefix(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldHasPrefix(FieldMbti, v))
}

// MbtiHasSuffix applies the HasSuffix predicate on the "mbti" field.
func MbtiHasSuffix(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldHasSuffix(FieldMbti, v))
}

// MbtiIsNil applies the IsNil predicate on the "mbti" field.
func MbtiIsNil() predicate.Avatar {
	return predicate.Avatar(sql.FieldIsNull(FieldMbti))
}

// MbtiNotNil applies the NotNil predicate on the "mbti" field.
func MbtiNotNil() predicate.Avatar {
	return predicate.Avatar(sql.FieldNotNull(FieldMbti))
}

// MbtiEqualFold applies the EqualFold predicate on the "mbti" field.
func MbtiEqualFold(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEqualFold(FieldMbti, v))
}

// MbtiContainsFold applies the ContainsFold predicate on the "mbti" field.
func MbtiContainsFold(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldContainsFold(FieldMbti, v))
}

// PictureEQ applies the EQ predicate on the "picture" field.
func PictureEQ(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldPicture, v))
}

// PictureNEQ applies the NEQ predicate on the "picture" field.
func PictureNEQ(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldNEQ(FieldPicture, v))
}

// PictureIn applies the In predicate on the "picture" field.
func PictureIn(vs ...string) predicate.Avatar {
	return predicate.Avatar(sql.FieldIn(FieldPicture, vs...))
}

// PictureNotIn applies the NotIn predicate on the "picture" field.
func PictureNotIn(vs ...string) predicate.Avatar {
	return predicate.Avatar(sql.FieldNotIn(FieldPicture, vs...))
}

// PictureGT applies the GT predicate on the "picture" field.
func PictureGT(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldGT(FieldPicture, v))
}

// PictureGTE applies the GTE predicate on the "picture" field.
func PictureGTE(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldGTE(FieldPicture, v))
}

// PictureLT applies the LT predicate on the "picture" field.
func PictureLT(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldLT(FieldPicture, v))
}

// PictureLTE applies the LTE predicate on the "picture" field.
func PictureLTE(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldLTE(FieldPicture, v))
}

// PictureContains applies the Contains predicate on the "picture" field.
func PictureContains(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldContains(FieldPicture, v))
}

// PictureHasPrefix applies the HasPrefix predicate on the "picture" field.
func PictureHasPrefix(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldHasPrefix(FieldPicture, v))
}

// PictureHasSuffix applies the HasSuffix predicate on the "picture" field.
func PictureHasSuffix(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldHasSuffix(FieldPicture, v))
}

// PictureEqualFold applies the EqualFold predicate on the "picture" field.
func PictureEqualFold(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEqualFold(FieldPicture, v))
}

// PictureContainsFold applies the ContainsFold predicate on the "picture" field.
func PictureContainsFold(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldContainsFold(FieldPicture, v))
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldNickname, v))
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldNEQ(FieldNickname, v))
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.Avatar {
	return predicate.Avatar(sql.FieldIn(FieldNickname, vs...))
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.Avatar {
	return predicate.Avatar(sql.FieldNotIn(FieldNickname, vs...))
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldGT(FieldNickname, v))
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldGTE(FieldNickname, v))
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldLT(FieldNickname, v))
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldLTE(FieldNickname, v))
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldContains(FieldNickname, v))
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldHasPrefix(FieldNickname, v))
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldHasSuffix(FieldNickname, v))
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEqualFold(FieldNickname, v))
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldContainsFold(FieldNickname, v))
}

// IntroduceEQ applies the EQ predicate on the "introduce" field.
func IntroduceEQ(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldIntroduce, v))
}

// IntroduceNEQ applies the NEQ predicate on the "introduce" field.
func IntroduceNEQ(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldNEQ(FieldIntroduce, v))
}

// IntroduceIn applies the In predicate on the "introduce" field.
func IntroduceIn(vs ...string) predicate.Avatar {
	return predicate.Avatar(sql.FieldIn(FieldIntroduce, vs...))
}

// IntroduceNotIn applies the NotIn predicate on the "introduce" field.
func IntroduceNotIn(vs ...string) predicate.Avatar {
	return predicate.Avatar(sql.FieldNotIn(FieldIntroduce, vs...))
}

// IntroduceGT applies the GT predicate on the "introduce" field.
func IntroduceGT(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldGT(FieldIntroduce, v))
}

// IntroduceGTE applies the GTE predicate on the "introduce" field.
func IntroduceGTE(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldGTE(FieldIntroduce, v))
}

// IntroduceLT applies the LT predicate on the "introduce" field.
func IntroduceLT(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldLT(FieldIntroduce, v))
}

// IntroduceLTE applies the LTE predicate on the "introduce" field.
func IntroduceLTE(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldLTE(FieldIntroduce, v))
}

// IntroduceContains applies the Contains predicate on the "introduce" field.
func IntroduceContains(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldContains(FieldIntroduce, v))
}

// IntroduceHasPrefix applies the HasPrefix predicate on the "introduce" field.
func IntroduceHasPrefix(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldHasPrefix(FieldIntroduce, v))
}

// IntroduceHasSuffix applies the HasSuffix predicate on the "introduce" field.
func IntroduceHasSuffix(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldHasSuffix(FieldIntroduce, v))
}

// IntroduceIsNil applies the IsNil predicate on the "introduce" field.
func IntroduceIsNil() predicate.Avatar {
	return predicate.Avatar(sql.FieldIsNull(FieldIntroduce))
}

// IntroduceNotNil applies the NotNil predicate on the "introduce" field.
func IntroduceNotNil() predicate.Avatar {
	return predicate.Avatar(sql.FieldNotNull(FieldIntroduce))
}

// IntroduceEqualFold applies the EqualFold predicate on the "introduce" field.
func IntroduceEqualFold(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldEqualFold(FieldIntroduce, v))
}

// IntroduceContainsFold applies the ContainsFold predicate on the "introduce" field.
func IntroduceContainsFold(v string) predicate.Avatar {
	return predicate.Avatar(sql.FieldContainsFold(FieldIntroduce, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.Avatar {
	return predicate.Avatar(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.Avatar {
	return predicate.Avatar(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.Avatar {
	return predicate.Avatar(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.Avatar {
	return predicate.Avatar(sql.FieldNotIn(FieldState, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Avatar {
	return predicate.Avatar(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Avatar {
	return predicate.Avatar(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAvatarRelations applies the HasEdge predicate on the "avatar_relations" edge.
func HasAvatarRelations() predicate.Avatar {
	return predicate.Avatar(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AvatarRelationsTable, AvatarRelationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAvatarRelationsWith applies the HasEdge predicate on the "avatar_relations" edge with a given conditions (other predicates).
func HasAvatarRelationsWith(preds ...predicate.Relation) predicate.Avatar {
	return predicate.Avatar(func(s *sql.Selector) {
		step := newAvatarRelationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFriendRelations applies the HasEdge predicate on the "friend_relations" edge.
func HasFriendRelations() predicate.Avatar {
	return predicate.Avatar(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FriendRelationsTable, FriendRelationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFriendRelationsWith applies the HasEdge predicate on the "friend_relations" edge with a given conditions (other predicates).
func HasFriendRelationsWith(preds ...predicate.Relation) predicate.Avatar {
	return predicate.Avatar(func(s *sql.Selector) {
		step := newFriendRelationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMessages applies the HasEdge predicate on the "messages" edge.
func HasMessages() predicate.Avatar {
	return predicate.Avatar(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMessagesWith applies the HasEdge predicate on the "messages" edge with a given conditions (other predicates).
func HasMessagesWith(preds ...predicate.Message) predicate.Avatar {
	return predicate.Avatar(func(s *sql.Selector) {
		step := newMessagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Avatar) predicate.Avatar {
	return predicate.Avatar(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Avatar) predicate.Avatar {
	return predicate.Avatar(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Avatar) predicate.Avatar {
	return predicate.Avatar(sql.NotPredicates(p))
}
