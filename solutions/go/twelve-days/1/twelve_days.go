package twelve

var number = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var present = []string{
    " twelve Drummers Drumming,",
    " eleven Pipers Piping,",
    " ten Lords-a-Leaping,",
    " nine Ladies Dancing,",
    " eight Maids-a-Milking,",
    " seven Swans-a-Swimming,",
    " six Geese-a-Laying,",
    " five Gold Rings,",
    " four Calling Birds,",
    " three French Hens,",
    " two Turtle Doves,",
    " a Partridge in a Pear Tree.",
}

func Verse(i int) string {
    result := "On the " + number[i-1] + " day of Christmas my true love gave to me:"
    if i == 1 {
        result += present[11]
    } else {
        for ;i > 0; i-- {
            if i == 1 {
                result += " and"
            }
            result += present[12-i]
        }
    }
	return result
}

func Song() string {
    result := ""
	for i := 1; i <= 12; i++ {
        result += Verse(i)
        if i < 12 {
        	result += "\n"
        }
    }
    return result
}
