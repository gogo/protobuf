package pyjsonmarshal

type PyClassNode struct {
	ClassName   string
	FieldDef    string
	ReflectDef  string
	ValidateDef string
}

const PyClassTpl = `
_{{.ClassName}}_reflect = {
{{.ReflectDef}}}

_{{.ClassName}}_validate = {
{{.ValidateDef}}}

class {{.ClassName}}(object):
{{.FieldDef}}

    @staticmethod
    def New():
        return {{.ClassName}}()

    def Marshal(self):
        if not self.Validate():
            logging.warning("marshal {{.ClassName}} failed")
            return False
        return json.dumps(self, default=lambda o: o.__dict__, sort_keys = True, indent = 4)

    def Unmarshal(self, data=None):
        if not isinstance(data, dict):
            try:
              data = json.loads(data)
            except ValueError as e:
                logging.warning("{}".format(e))
                return False

        for key in data:
            subdata = data[key]
            if key in _{{.ClassName}}_reflect:
                if type(subdata) is list:
                    array = []
                    for v in subdata:
                        val = _{{.ClassName}}_reflect[key].New()
                        if not val.Unmarshal(data[key]):
                            logging.warning("marshal key={} object error={}".format(key))
                            return False
                        array.append(val)
                    self.__dict__[key] = array
                else:
                    val = _{{.ClassName}}_reflect[key].New()
                    if not val.Unmarshal(data[key]):
                        logging.warning("marshal key={} object error={}".format(key))
                        return False
                    self.__dict__[key] = val
            else:
                self.__dict__[key] = subdata
        return self.Validate()

    def Validate(self):
        for key in _{{.ClassName}}_validate:
            if key not in self.__dict__:
                logging.warning("required key={} is None".format(key))
                return False
            if key in _{{.ClassName}}_reflect and self.__dict__[key].Validate():
                return False

        return True

`
