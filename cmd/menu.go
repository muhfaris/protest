package cmd

import "log"

func screen() {
	log.Println(`
	______          _            _   
	| ___ \        | |          | |  
	| |_/ / __ ___ | |_ ___  ___| |_ 
	|  __/ '__/ _ \| __/ _ \/ __| __|
	| |  | | | (_) | ||  __/\__ \ |_ 
	\_|  |_|  \___/ \__\___||___/\__|
									 
			-Producer Testing`)
}

// --------------------------------------------------------
// | - MENU
// | --> send new data based on template (new)
// | --> update existing data (update)
// | --> change data operation to insert (i) / update (u)
// | --> send data forever with interval time (interval)
// --------------------------------------------------------

func menu() {
}
