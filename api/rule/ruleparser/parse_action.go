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
	/*
		condition=
		type,val,val,const,const
	*/
	actionList := strings.Split(action, ";")
	for _, actionStr := range actionList {
		// actionType_01:P_01($val01,$val_02)
		var aL []string
		for i, c := range actionStr {
			if string(c) == ":" {
				if i == len(actionStr)-1 {
					aL = append(aL, actionStr)
					aL = append(aL, "")
				} else {
					aL = append(aL, actionStr[:i])
					aL = append(aL, actionStr[i+1:])
				}
			}
		}

		aL[0] = strings.Replace(aL[0], " ", "", -1)
		result = append(result, Action{
			ActionType:   aL[0],
			ActionParams: aL[1],
		})
	}
	return result
}
