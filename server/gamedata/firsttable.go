package gamedata 

type FirstTable struct {
    ID int
    F_Name string
    Age int
    LevelAPValue []int

}
var (
    FirstTableData = make(map[int]FirstTable)
)

func  FirstTableinit() {
	rf := readRf(FirstTable{})
	for i := 0; i < rf.NumRecord(); i++ {
		r := rf.Record(i).(*FirstTable)
        FirstTableData[r.ID] = *r
    }
}

func GetFirstTableByID(id int) (FirstTable) {
	return  FirstTableData[id]
}
