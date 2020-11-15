package main

import (
	"fmt"
	student ".."
)

func main() {
	str := "HAHelloHAhowHAareHAyou?"
	fmt.Println(student.Split(str, "HA"))

	str2 := "aaa man a plan a canal panama"
	fmt.Printf("%q\n", student.Split(str2, "a"))

	str3 := "a,b,c"
	fmt.Printf("%q\n", student.Split(str3, ","))

	str4 := " xyz "
	fmt.Printf("%q\n", student.Split(str4, ""))

	str5 := ""
	fmt.Printf("%q\n", student.Split(str5, "Bernardo O'Higgins"))

	fmt.Println()
	fmt.Println()

	str6 := "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,"
	fmt.Println(student.Split(str6, "|=choumi=|"))
	fmt.Println()

	str7 := "!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,"
	fmt.Println(student.Split(str7, "!==!"))
	fmt.Println()

	str8 := "AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,"
	fmt.Println(student.Split(str8, "AFJ"))
	fmt.Println()

	str9 := "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"
	fmt.Println(student.Split(str9, "<<==123==>>"))
	fmt.Println()

}
