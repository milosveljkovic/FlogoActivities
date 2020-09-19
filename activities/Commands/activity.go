package sample

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMetadata  = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debugf("Setting: %s", s.ASetting)

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMetadata 
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	strCommands := ""

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	if(input.Lux >=150){ 			 				//DAY
		if (input.Lux >= 200 ) { 				//CRITICAL
			strCommands += "/TOD1" 
		}else {
			strCommands += "/FOD1"
		}

		if(input.Temperature >= 30){ 		//CRITICAL
			strCommands += "/TOD2"
		}else {
			strCommands += "/FOD2"
		}

		if(input.Humidity >= 50){ 			//CRITICAL
			strCommands += "/TOD3"
		}else {
			strCommands += "/FOD3"
		}
	}else {														//NIGHT
		strCommands += "/FOD1"				  //NO SIGNAL					
		if(input.Temperature <= 20){ 		//CRITICAL
			strCommands += "/TOD2"
		}else {
			strCommands += "/FOD2"
		}
		if(input.Humidity >= 50){ 			//CRITICAL
			strCommands += "/TOD3"
		}else {
			strCommands += "/FOD3"
		}
	}



	output := &Output{
					Commands: strCommands,
				}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
