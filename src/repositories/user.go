package repositories

// ... спизж.ено полностью
import (
	"buridansAss/chat/src/models"
	"github.com/go-pg/pg"
	"time"
)

type (
	User interface {
		Create(name, email, password string) (*models.User, error)
		FindByEmail(email string) (*models.User, error)
		FindByToken(token string) (*models.User, error)
		UpdateToken(user *models.User, token string) error
	}

	userRepository struct {
		db *pg.DB
	}
)

func NewUser(db *pg.DB) User {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(name, email, pass string) (*models.User, error) {
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: pass,
	}

	if _, err := u.db.Model(&user).Where("email=?", email).SelectOrInsert(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) findBy(condition string, val interface{}) (*models.User, error) {
	var user models.User

	if err := u.db.Model(&user).Where(condition, val).First(); err != nil {
		if err != nil {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (u *userRepository) FindByEmail(email string) (*models.User, error) {
	return u.findBy("email=?", email)
}

func (u *userRepository) FindByToken(token string) (*models.User, error) {
	return u.findBy("token=?", token)
}

func (u *userRepository) UpdateToken(user *models.User, token string) error {
	user.Token = token
	user.UpdatedAt = time.Now()

	if _, err := u.db.Model(user).Column("token", "updated_at").WherePK().Update(); err != nil {
		return err
	}

	return nil
}
