package ruleparser

import "strings"

type Action struct {
	ActionType   string
	ActionParams string
}

/*
	Action=
	actionType_01: P_01($val01, $val_02); actionType_02: P_02($val_01, $val_02)
*/
func ParseAction(action string) []Action {

	var result []Action
	action = strings.Replace(action, " ", "", -1)
	/*
		condition=
		type,val,val,const,const
	*/
	actionList := strings.Split(action, ";")
	for _, actionStr := range actionList {
		// actionType_01:P_01($val01,$val_02)
		aL := strings.Split(actionStr, ":")
		result = append(result, Action{
			ActionType:   aL[0],
			ActionParams: aL[1],
		})
	}
	return result
}
