/**
 * @author		Ian M. Fink
 *
 * @file		main.go
 *
 * Copyright (C) 2025 Ian M. Fink.  All rights reserved.
 *
 * This program is free software:  you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by the Free
 * Software Foundation, either version 3 of the License, or (at your option)
 * any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
 * or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
 * more details.
 *
 * You should have received a copy of the GNU General Public License along
 * with this program.  If not, please see: https://www.gnu.org/licenses.
 *
 * Tabstop:	4
 *
 */

package main

/*
 * Imports
 */

import (
	"fmt"
	"strconv"
	"time"
	"strings"
	"path"
	"os"
)

/**********************************************************************/

func main() {
	var (
		theTime		time.Time
		tmpInt64	int64
		theYear		int
		theMonth	time.Month
		theDay		int
		daysToAdd	int
		dateSplit	[]string
		err			error
	)

	if len(os.Args) < 5 {
		fmt.Printf("Usage: %s <year> <month #> <day> <days to add>\n",
			path.Base(os.Args[0]))
		os.Exit(10)
	}

	tmpInt64, err = strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	theYear = int(tmpInt64)

	tmpInt64, err = strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	theMonth = time.Month(tmpInt64)

	tmpInt64, err = strconv.ParseInt(os.Args[3], 10, 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	theDay = int(tmpInt64)

	tmpInt64, err = strconv.ParseInt(os.Args[4], 10, 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
	daysToAdd = int(tmpInt64)

//	func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	theTime = time.Date(theYear, theMonth, theDay,
		0,	// hour
		0,	// min
		0,	// sec
		0,	// nsec
		time.UTC)	// *time.Location

	dateSplit = strings.Split(theTime.String(), " ")
	fmt.Printf("%s plus %d days is ", dateSplit[0], daysToAdd)
	dateSplit = strings.Split(theTime.Add(time.Hour * 24 *
		time.Duration(daysToAdd)).String(), " ")
	fmt.Printf("%s\n", dateSplit[0])

} /* main */

/*
 * End of file:	main.go
 */

