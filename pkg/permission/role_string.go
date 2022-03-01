// Code generated by "stringer -type=Role -linecomment -trimprefix Role"; DO NOT EDIT.

package permission

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RoleUser-0]
	_ = x[RoleSiteAdmin-5]
	_ = x[RoleService-6]
}

const (
	_Role_name_0 = "USER"
	_Role_name_1 = "SITE_ADMINSERVICE"
)

var (
	_Role_index_1 = [...]uint8{0, 10, 17}
)

func (i Role) String() string {
	switch {
	case i == 0:
		return _Role_name_0
	case 5 <= i && i <= 6:
		i -= 5
		return _Role_name_1[_Role_index_1[i]:_Role_index_1[i+1]]
	default:
		return "Role(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
