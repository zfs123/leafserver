package gamedata 

type SecondTable struct {
    ID int
    F_Name string
    Age int
    LevelAPValue []int

}
var (
    SecondTableData = make(map[int]SecondTable)
)

func  SecondTableinit() {
	rf := readRf(SecondTable{})
	for i := 0; i < rf.NumRecord(); i++ {
		r := rf.Record(i).(*SecondTable)
        SecondTableData[r.ID] = *r
    }
}

func GetSecondTableByID(id int) (SecondTable) {
	return  SecondTableData[id]
}
