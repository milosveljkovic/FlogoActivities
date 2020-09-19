package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASetting string `md:"aSetting,required"`
}

type Input struct {
	Lux int `md:"lux,required"`

	Humidity int `md:"humidity,required"`

	Temperature int `md:"temperature,required"`
}

// FromMap converts the values from a map into the struct Input
func (i *Input) FromMap(values map[string]interface{}) error {
	lux, err := coerce.ToInt(values["lux"])
	if err != nil {
		return err
	}
	i.Lux = lux

	humidity, err := coerce.ToInt(values["humidity"])
	if err != nil {
		return err
	}
	i.Humidity = humidity

	temperature, err := coerce.ToInt(values["temperature"])
	if err != nil {
		return err
	}
	i.Temperature = temperature

	return nil
}

// ToMap converts the struct Input into a map
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"lux":   i.Lux,
		"humidity": i.Humidity,
		"temperature":  i.Temperature,
	}
}

type Output struct {
	Commands string `md:"commands"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["commands"])
	o.Commands = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"commands": o.Commands,
	}
}
