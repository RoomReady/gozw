package functions

type GetInitData struct {
	FunctionId uint8
}

func NewGetInitData(nodeId uint8) GetInitData {
	return GetInitData{
		FunctionId: ZwGetInitData,
	}
}

func (f *GetInitData) Marshal() []byte {
	return []byte{
		f.FunctionId,
	}
}
