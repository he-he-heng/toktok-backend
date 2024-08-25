// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"
	"toktok-backend/internal/adapter/persistence/mysql/ent/avatar"
	"toktok-backend/internal/adapter/persistence/mysql/ent/message"
	"toktok-backend/internal/adapter/persistence/mysql/ent/relation"
	"toktok-backend/internal/adapter/persistence/mysql/ent/schema"
	"toktok-backend/internal/adapter/persistence/mysql/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	avatarMixin := schema.Avatar{}.Mixin()
	avatarMixinHooks0 := avatarMixin[0].Hooks()
	avatar.Hooks[0] = avatarMixinHooks0[0]
	avatarMixinInters0 := avatarMixin[0].Interceptors()
	avatar.Interceptors[0] = avatarMixinInters0[0]
	avatarMixinFields1 := avatarMixin[1].Fields()
	_ = avatarMixinFields1
	avatarFields := schema.Avatar{}.Fields()
	_ = avatarFields
	// avatarDescCreatedAt is the schema descriptor for created_at field.
	avatarDescCreatedAt := avatarMixinFields1[0].Descriptor()
	// avatar.DefaultCreatedAt holds the default value on creation for the created_at field.
	avatar.DefaultCreatedAt = avatarDescCreatedAt.Default.(func() time.Time)
	// avatarDescUpdatedAt is the schema descriptor for updated_at field.
	avatarDescUpdatedAt := avatarMixinFields1[1].Descriptor()
	// avatar.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	avatar.DefaultUpdatedAt = avatarDescUpdatedAt.Default.(func() time.Time)
	// avatar.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	avatar.UpdateDefaultUpdatedAt = avatarDescUpdatedAt.UpdateDefault.(func() time.Time)
	// avatarDescBirthday is the schema descriptor for birthday field.
	avatarDescBirthday := avatarFields[1].Descriptor()
	// avatar.BirthdayValidator is a validator for the "birthday" field. It is called by the builders before save.
	avatar.BirthdayValidator = func() func(string) error {
		validators := avatarDescBirthday.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(birthday string) error {
			for _, fn := range fns {
				if err := fn(birthday); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// avatarDescMbti is the schema descriptor for mbti field.
	avatarDescMbti := avatarFields[2].Descriptor()
	// avatar.MbtiValidator is a validator for the "mbti" field. It is called by the builders before save.
	avatar.MbtiValidator = func() func(string) error {
		validators := avatarDescMbti.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(mbti string) error {
			for _, fn := range fns {
				if err := fn(mbti); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// avatarDescNickname is the schema descriptor for nickname field.
	avatarDescNickname := avatarFields[4].Descriptor()
	// avatar.NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	avatar.NicknameValidator = avatarDescNickname.Validators[0].(func(string) error)
	messageMixin := schema.Message{}.Mixin()
	messageMixinHooks0 := messageMixin[0].Hooks()
	message.Hooks[0] = messageMixinHooks0[0]
	messageMixinInters0 := messageMixin[0].Interceptors()
	message.Interceptors[0] = messageMixinInters0[0]
	messageMixinFields1 := messageMixin[1].Fields()
	_ = messageMixinFields1
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreatedAt is the schema descriptor for created_at field.
	messageDescCreatedAt := messageMixinFields1[0].Descriptor()
	// message.DefaultCreatedAt holds the default value on creation for the created_at field.
	message.DefaultCreatedAt = messageDescCreatedAt.Default.(func() time.Time)
	// messageDescUpdatedAt is the schema descriptor for updated_at field.
	messageDescUpdatedAt := messageMixinFields1[1].Descriptor()
	// message.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	message.DefaultUpdatedAt = messageDescUpdatedAt.Default.(func() time.Time)
	// message.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	message.UpdateDefaultUpdatedAt = messageDescUpdatedAt.UpdateDefault.(func() time.Time)
	// messageDescContent is the schema descriptor for content field.
	messageDescContent := messageFields[1].Descriptor()
	// message.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	message.ContentValidator = messageDescContent.Validators[0].(func(string) error)
	// messageDescEnteredAt is the schema descriptor for entered_at field.
	messageDescEnteredAt := messageFields[2].Descriptor()
	// message.DefaultEnteredAt holds the default value on creation for the entered_at field.
	message.DefaultEnteredAt = messageDescEnteredAt.Default.(func() time.Time)
	relationMixin := schema.Relation{}.Mixin()
	relationMixinHooks0 := relationMixin[0].Hooks()
	relation.Hooks[0] = relationMixinHooks0[0]
	relationMixinInters0 := relationMixin[0].Interceptors()
	relation.Interceptors[0] = relationMixinInters0[0]
	relationMixinFields1 := relationMixin[1].Fields()
	_ = relationMixinFields1
	relationFields := schema.Relation{}.Fields()
	_ = relationFields
	// relationDescCreatedAt is the schema descriptor for created_at field.
	relationDescCreatedAt := relationMixinFields1[0].Descriptor()
	// relation.DefaultCreatedAt holds the default value on creation for the created_at field.
	relation.DefaultCreatedAt = relationDescCreatedAt.Default.(func() time.Time)
	// relationDescUpdatedAt is the schema descriptor for updated_at field.
	relationDescUpdatedAt := relationMixinFields1[1].Descriptor()
	// relation.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	relation.DefaultUpdatedAt = relationDescUpdatedAt.Default.(func() time.Time)
	// relation.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	relation.UpdateDefaultUpdatedAt = relationDescUpdatedAt.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	userMixinInters0 := userMixin[0].Interceptors()
	user.Interceptors[0] = userMixinInters0[0]
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields1[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields1[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescUID is the schema descriptor for uid field.
	userDescUID := userFields[0].Descriptor()
	// user.UIDValidator is a validator for the "uid" field. It is called by the builders before save.
	user.UIDValidator = func() func(string) error {
		validators := userDescUID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(uid string) error {
			for _, fn := range fns {
				if err := fn(uid); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
}

const (
	Version = "v0.14.0"                                         // Version of ent codegen.
	Sum     = "h1:EO3Z9aZ5bXJatJeGqu/EVdnNr6K4mRq3rWe5owt0MC4=" // Sum of ent codegen.
)
