package kernel

import (
	"context"
	pb "control/grpc/control"
	sq "github.com/Masterminds/squirrel"
	"log"
)

const addUserToStaffTable = `
INSERT INTO staff
DEFAULT VALUES
RETURNING id;
`

func (s *Kernel) AddNewUserOrReturn(ctx context.Context) (int, error) {
	rows, err := s.db.QueryContext(ctx, addUserToStaffTable)
	if err != nil {
		return 0, err
	}
	if !rows.Next() {
		panic("No created element")
	}

	var userId int
	_ = rows.Scan(&userId)

	return userId, nil
}

func (s *Kernel) UpdateShopCommonPart(ctx context.Context, shopId int, info *pb.OptionalBotInfo) error {
	commonInfo := info.GetCommonBotInfo()
	if commonInfo == nil {
		return nil
	}

	shouldUpdate := false
	builder := sq.Update("shop").Where(sq.Eq{"shop_id": shopId})
	if commonInfo.Name != nil {
		builder = builder.Set("name", *commonInfo.Name)
		shouldUpdate = true
	}

	if shouldUpdate {
		queue, _, _ := builder.ToSql()
		_, err := s.db.ExecContext(ctx, queue)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Kernel) DeleteShopCommonPart(ctx context.Context, shopId int) error {
	whereCond := sq.Eq{"shop_id": shopId}

	queue, _, _ := sq.Delete("shop").Where(whereCond).ToSql()
	_, err := s.db.ExecContext(ctx, queue)
	if err != nil {
		return err
	}

	queue, _, _ = sq.Delete("shop_staff_roles").Where(whereCond).ToSql()
	_, err = s.db.ExecContext(ctx, queue)
	if err != nil {
		return err
	}

	queue, _, _ = sq.Delete("vk_client_group").Where(whereCond).ToSql()
	_, err = s.db.ExecContext(ctx, queue)
	if err != nil {
		return err
	}

	return nil
}

func (s *Kernel) GetShopIdByName(ctx context.Context, name string) (int, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id FROM shop WHERE name = $1", name)
	if err != nil {
		return 0, err
	}
	if !rows.Next() {
		return 0, &NoSuchItemError{}
	}

	var shopId int
	_ = rows.Scan(&shopId)
	return shopId, nil
}

func (s *Kernel) GetShopIdByVkGroupId(ctx context.Context, groupId int) (int, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT shop_id FROM vk_client_group WHERE vk_group_id = $1", groupId)
	if err != nil {
		return 0, err
	}
	if !rows.Next() {
		return 0, &NoSuchItemError{}
	}

	var shopId int
	_ = rows.Scan(&shopId)
	return shopId, nil
}

func (s *Kernel) GetShopIdByKey(ctx context.Context, key *pb.CommonShopKey) (int, error) {
	switch op := key.CommonKey.(type) {
	case *pb.CommonShopKey_Id:
		return int(op.Id), nil
	case *pb.CommonShopKey_Name:
		return s.GetShopIdByName(ctx, op.Name)
	case *pb.CommonShopKey_VkGroupId:
		return s.GetShopIdByVkGroupId(ctx, int(op.VkGroupId))
	default:
		log.Fatal("adf")
	}
	return 0, nil
}

func (s *Kernel) CheckForRole(ctx context.Context, shopId int, userId int, role string) (bool, error) {
	roles := sq.Or{sq.Eq{"role": "creator"}}
	if role != "creator" {
		roles = append(roles, sq.Eq{"role": "admin"})
	}
	if role != "admin" {
		roles = append(roles, sq.Eq{"role": role})
	}
	queue, args, _ := sq.Select("1").From("shop_staff_roles").Where(sq.And{
		sq.Eq{"shop_id": shopId},
		sq.Eq{"user_id": userId},
		roles,
	}).PlaceholderFormat(sq.Dollar).ToSql()

	rows, err := s.db.QueryContext(ctx, queue, args...)
	defer rows.Close()
	if err != nil {
		return false, err
	}

	exists := rows.Next()
	return exists, nil
}

func (s *Kernel) AddRole(ctx context.Context, shopId int, userId int, role string) error {
	queue, _, _ := sq.Insert("shop_staff_roles").Columns("shop_id", "user_id", "role").
		Values(shopId, userId, role).ToSql()

	_, err := s.db.ExecContext(ctx, queue)
	if err != nil {
		return err
	}

	return nil
}
