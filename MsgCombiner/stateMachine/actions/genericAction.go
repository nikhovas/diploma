package actions

type Arguments map[string]string
type Returns map[string]string

type GenericAction struct {
	Name      string
	Arguments Arguments
	Return    Returns
}
