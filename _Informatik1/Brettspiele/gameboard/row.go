package gameboard

type Row []string


func (row Row) ContainsOnly(symbol string) bool {

for _,character := range row {
	if character != symbol {
		return false
	}
}
return true

}
