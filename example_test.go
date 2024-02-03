package vegagoja_test

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/xo/vegagoja"
)

func Example() {
	vega := vegagoja.New(
		vegagoja.WithPrefixedSourceDir("data/", "testdata/data/"),
	)
	svg, err := vega.Render(context.Background(), candlestickSpec)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("candlestick.svg", []byte(svg), 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(svg))
	// Output:
	// <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" class="marks" width="448" height="273" viewBox="0 0 448 273"><rect width="448" height="273" fill="white"/><g fill="none" stroke-miterlimit="10" transform="translate(43,10)"><g class="mark-group role-frame root" role="graphics-object" aria-roledescription="group mark container"><g transform="translate(0,0)"><path class="background" aria-hidden="true" d="M0.5,0.5h400v200h-400Z" stroke="#ddd"/><g><g class="mark-group role-axis" aria-hidden="true"><g transform="translate(0.5,200.5)"><path class="background" aria-hidden="true" d="M0,0h0v0h0Z" pointer-events="none"/><g><g class="mark-rule role-axis-grid" pointer-events="none"><line transform="translate(42,-200)" x2="0" y2="200" stroke="#ddd" stroke-width="1" opacity="1"/><line transform="translate(88,-200)" x2="0" y2="200" stroke="#ddd" stroke-width="1" opacity="1"/><line transform="translate(133,-200)" x2="0" y2="200" stroke="#ddd" stroke-width="1" opacity="1"/><line transform="translate(179,-200)" x2="0" y2="200" stroke="#ddd" stroke-width="1" opacity="1"/><line transform="translate(224,-200)" x2="0" y2="200" stroke="#ddd" stroke-width="1" opacity="1"/><line transform="translate(270,-200)" x2="0" y2="200" stroke="#ddd" stroke-width="1" opacity="1"/><line transform="translate(315,-200)" x2="0" y2="200" stroke="#ddd" stroke-width="1" opacity="1"/><line transform="translate(361,-200)" x2="0" y2="200" stroke="#ddd" stroke-width="1" opacity="1"/></g></g><path class="foreground" aria-hidden="true" d="" pointer-events="none" display="none"/></g></g><g class="mark-group role-axis" aria-hidden="true"><g transform="translate(0.5,0.5)"><path class="background" aria-hidden="true" d="M0,0h0v0h0Z" pointer-events="none"/><g><g class="mark-rule role-axis-grid" pointer-events="none"><line transform="translate(0,182)" x2="400" y2="0" stroke="#ddd" stroke-width="1" opacity="1"/><line transform="translate(0,145)" x2="400" y2="0" stroke="#ddd" stroke-width="1" opacity="1"/><line transform="translate(0,109)" x2="400" y2="0" stroke="#ddd" stroke-width="1" opacity="1"/><line transform="translate(0,73)" x2="400" y2="0" stroke="#ddd" stroke-width="1" opacity="1"/><line transform="translate(0,36)" x2="400" y2="0" stroke="#ddd" stroke-width="1" opacity="1"/><line transform="translate(0,0)" x2="400" y2="0" stroke="#ddd" stroke-width="1" opacity="1"/></g></g><path class="foreground" aria-hidden="true" d="" pointer-events="none" display="none"/></g></g><g class="mark-group role-axis" role="graphics-symbol" aria-roledescription="axis" aria-label="X-axis titled 'Date in 2009' for a time scale with values from 05/31 to 08/01"><g transform="translate(0.5,200.5)"><path class="background" aria-hidden="true" d="M0,0h0v0h0Z" pointer-events="none"/><g><g class="mark-rule role-axis-tick" pointer-events="none"><line transform="translate(42,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(88,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(133,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(179,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(224,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(270,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(315,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(361,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/></g><g class="mark-text role-axis-label" pointer-events="none"><text text-anchor="end" transform="translate(42.10416670656678,7) rotate(315) translate(0,8)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">06/07</text><text text-anchor="end" transform="translate(87.6041667059085,7) rotate(315) translate(0,8)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">06/14</text><text text-anchor="end" transform="translate(133.10416670525024,7) rotate(315) translate(0,8)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">06/21</text><text text-anchor="end" transform="translate(178.60416670459196,7) rotate(315) translate(0,8)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">06/28</text><text text-anchor="end" transform="translate(224.1041667039337,7) rotate(315) translate(0,8)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">07/05</text><text text-anchor="end" transform="translate(269.6041667032754,7) rotate(315) translate(0,8)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">07/12</text><text text-anchor="end" transform="translate(315.10416670261714,7) rotate(315) translate(0,8)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">07/19</text><text text-anchor="end" transform="translate(360.6041667019589,7) rotate(315) translate(0,8)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">07/26</text></g><g class="mark-rule role-axis-domain" pointer-events="none"><line transform="translate(0,0)" x2="400" y2="0" stroke="#888" stroke-width="1" opacity="1"/></g><g class="mark-text role-axis-title" pointer-events="none"><text text-anchor="middle" transform="translate(200,55.355339059327406)" font-family="sans-serif" font-size="11px" font-weight="bold" fill="#000" opacity="1">Date in 2009</text></g></g><path class="foreground" aria-hidden="true" d="" pointer-events="none" display="none"/></g></g><g class="mark-group role-axis" role="graphics-symbol" aria-roledescription="axis" aria-label="Y-axis titled 'Price' for a linear scale with values from 23 to 34"><g transform="translate(0.5,0.5)"><path class="background" aria-hidden="true" d="M0,0h0v0h0Z" pointer-events="none"/><g><g class="mark-rule role-axis-tick" pointer-events="none"><line transform="translate(0,182)" x2="-5" y2="0" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(0,145)" x2="-5" y2="0" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(0,109)" x2="-5" y2="0" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(0,73)" x2="-5" y2="0" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(0,36)" x2="-5" y2="0" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(0,0)" x2="-5" y2="0" stroke="#888" stroke-width="1" opacity="1"/></g><g class="mark-text role-axis-label" pointer-events="none"><text text-anchor="end" transform="translate(-7,184.8181818181818)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">24</text><text text-anchor="end" transform="translate(-7,148.45454545454547)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">26</text><text text-anchor="end" transform="translate(-7,112.09090909090908)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">28</text><text text-anchor="end" transform="translate(-7,75.72727272727273)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">30</text><text text-anchor="end" transform="translate(-7,39.36363636363635)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">32</text><text text-anchor="end" transform="translate(-7,3)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">34</text></g><g class="mark-rule role-axis-domain" pointer-events="none"><line transform="translate(0,200)" x2="0" y2="-200" stroke="#888" stroke-width="1" opacity="1"/></g><g class="mark-text role-axis-title" pointer-events="none"><text text-anchor="middle" transform="translate(-27,100) rotate(-90) translate(0,-2)" font-family="sans-serif" font-size="11px" font-weight="bold" fill="#000" opacity="1">Price</text></g></g><path class="foreground" aria-hidden="true" d="" pointer-events="none" display="none"/></g></g><g class="mark-rule role-mark layer_0_marks" role="graphics-object" aria-roledescription="rule mark container"><line aria-label="Date in 2009: 06/01; low: 28.45; high: 30.05" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(5.000000040436921,100.90909090909092)" x2="0" y2="-29.090909090909122" stroke="#06982d"/><line aria-label="Date in 2009: 06/02; low: 28.3; high: 30.13" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(11.500000040342883,103.63636363636361)" x2="0" y2="-33.272727272727224" stroke="#ae1325"/><line aria-label="Date in 2009: 06/03; low: 29.62; high: 31.79" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(18.000000040248842,79.63636363636361)" x2="0" y2="-39.4545454545454" stroke="#06982d"/><line aria-label="Date in 2009: 06/04; low: 29.92; high: 31.02" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(24.500000040154806,74.18181818181814)" x2="0" y2="-19.999999999999943" stroke="#ae1325"/><line aria-label="Date in 2009: 06/05; low: 28.85; high: 30.81" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(31.000000040060765,93.63636363636361)" x2="0" y2="-35.63636363636358" stroke="#06982d"/><line aria-label="Date in 2009: 06/08; low: 26.41; high: 31.82" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(50.50000003977865,138)" x2="0" y2="-98.36363636363637" stroke="#ae1325"/><line aria-label="Date in 2009: 06/09; low: 27.79; high: 29.77" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(57.00000003968461,112.90909090909093)" x2="0" y2="-36.000000000000014" stroke="#ae1325"/><line aria-label="Date in 2009: 06/10; low: 26.9; high: 29.74" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(63.50000003959057,129.0909090909091)" x2="0" y2="-51.63636363636361" stroke="#06982d"/><line aria-label="Date in 2009: 06/11; low: 26.81; high: 28.11" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(70.00000003949653,130.72727272727275)" x2="0" y2="-23.63636363636364" stroke="#06982d"/><line aria-label="Date in 2009: 06/12; low: 27.73; high: 28.5" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(76.5000000394025,113.99999999999999)" x2="0" y2="-13.999999999999986" stroke="#06982d"/><line aria-label="Date in 2009: 06/15; low: 29.64; high: 31.09" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(96.00000003912037,79.27272727272727)" x2="0" y2="-26.36363636363636" stroke="#06982d"/><line aria-label="Date in 2009: 06/16; low: 30.07; high: 32.75" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(102.50000003902633,71.45454545454545)" x2="0" y2="-48.72727272727272" stroke="#06982d"/><line aria-label="Date in 2009: 06/17; low: 30.64; high: 32.77" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(109.00000003893228,61.090909090909086)" x2="0" y2="-38.72727272727279" stroke="#06982d"/><line aria-label="Date in 2009: 06/18; low: 29.6; high: 31.54" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(115.50000003883825,79.99999999999999)" x2="0" y2="-35.27272727272724" stroke="#ae1325"/><line aria-label="Date in 2009: 06/19; low: 27.56; high: 29.32" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(122.0000000387442,117.09090909090911)" x2="0" y2="-32.00000000000003" stroke="#ae1325"/><line aria-label="Date in 2009: 06/22; low: 30.3; high: 32.05" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(141.5000000384621,67.27272727272727)" x2="0" y2="-31.81818181818175" stroke="#06982d"/><line aria-label="Date in 2009: 06/23; low: 27.83; high: 31.54" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(148.00000003836806,112.1818181818182)" x2="0" y2="-67.45454545454545" stroke="#ae1325"/><line aria-label="Date in 2009: 06/24; low: 28.79; high: 30.58" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(154.500000038274,94.72727272727275)" x2="0" y2="-32.54545454545455" stroke="#ae1325"/><line aria-label="Date in 2009: 06/25; low: 26.3; high: 29.56" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(161.00000003818,140)" x2="0" y2="-59.27272727272724" stroke="#ae1325"/><line aria-label="Date in 2009: 06/26; low: 25.76; high: 27.22" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(167.50000003808594,149.8181818181818)" x2="0" y2="-26.54545454545452" stroke="#ae1325"/><line aria-label="Date in 2009: 06/29; low: 25.29; high: 27.18" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(187.00000003780383,158.36363636363637)" x2="0" y2="-34.363636363636346" stroke="#ae1325"/><line aria-label="Date in 2009: 06/30; low: 25.02; high: 27.38" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(193.50000003770978,163.27272727272728)" x2="0" y2="-42.90909090909089" stroke="#06982d"/><line aria-label="Date in 2009: 07/01; low: 24.8; high: 26.31" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(200.00000003761573,167.27272727272725)" x2="0" y2="-27.45454545454541" stroke="#06982d"/><line aria-label="Date in 2009: 07/02; low: 26.22; high: 28.62" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(206.50000003752172,141.45454545454547)" x2="0" y2="-43.63636363636367" stroke="#06982d"/><line aria-label="Date in 2009: 07/06; low: 28.99; high: 30.6" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(232.50000003714555,91.09090909090911)" x2="0" y2="-29.272727272727316" stroke="#ae1325"/><line aria-label="Date in 2009: 07/07; low: 28.9; high: 30.94" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(239.00000003705148,92.72727272727275)" x2="0" y2="-37.09090909090913" stroke="#06982d"/><line aria-label="Date in 2009: 07/08; low: 30.43; high: 33.05" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(245.50000003695746,64.90909090909092)" x2="0" y2="-47.63636363636359" stroke="#06982d"/><line aria-label="Date in 2009: 07/09; low: 29.28; high: 30.49" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(252.00000003686344,85.8181818181818)" x2="0" y2="-21.999999999999957" stroke="#ae1325"/><line aria-label="Date in 2009: 07/10; low: 28.82; high: 30.34" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(258.5000000367694,94.18181818181817)" x2="0" y2="-27.636363636363626" stroke="#ae1325"/><line aria-label="Date in 2009: 07/13; low: 25.42; high: 29.24" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(278.0000000364873,155.99999999999997)" x2="0" y2="-69.4545454545454" stroke="#ae1325"/><line aria-label="Date in 2009: 07/14; low: 24.99; high: 26.84" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(284.5000000363932,163.81818181818184)" x2="0" y2="-33.636363636363654" stroke="#ae1325"/><line aria-label="Date in 2009: 07/15; low: 23.83; high: 26.06" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(291.0000000362992,184.90909090909093)" x2="0" y2="-40.54545454545456" stroke="#06982d"/><line aria-label="Date in 2009: 07/16; low: 24.51; high: 26.18" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(297.50000003620517,172.5454545454545)" x2="0" y2="-30.363636363636317" stroke="#ae1325"/><line aria-label="Date in 2009: 07/17; low: 23.88; high: 25.55" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(304.00000003611115,184.00000000000003)" x2="0" y2="-30.363636363636402" stroke="#ae1325"/><line aria-label="Date in 2009: 07/20; low: 24.26; high: 25.42" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(323.500000035829,177.09090909090907)" x2="0" y2="-21.090909090909093" stroke="#ae1325"/><line aria-label="Date in 2009: 07/21; low: 23.81; high: 25.14" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(330.00000003573496,185.2727272727273)" x2="0" y2="-24.181818181818215" stroke="#ae1325"/><line aria-label="Date in 2009: 07/22; low: 23.24; high: 24.14" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(336.50000003564094,195.63636363636368)" x2="0" y2="-16.36363636363643" stroke="#ae1325"/><line aria-label="Date in 2009: 07/23; low: 23.21; high: 24.05" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(343.00000003554686,196.18181818181816)" x2="0" y2="-15.272727272727252" stroke="#ae1325"/><line aria-label="Date in 2009: 07/24; low: 23; high: 23.87" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(349.50000003545284,200)" x2="0" y2="-15.818181818181841" stroke="#ae1325"/><line aria-label="Date in 2009: 07/27; low: 24.02; high: 24.86" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(369.00000003517073,181.45454545454547)" x2="0" y2="-15.27272727272728" stroke="#06982d"/><line aria-label="Date in 2009: 07/28; low: 24.28; high: 25.61" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(375.5000000350767,176.72727272727272)" x2="0" y2="-24.18181818181816" stroke="#06982d"/><line aria-label="Date in 2009: 07/29; low: 25.41; high: 26.18" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(382.00000003498263,156.1818181818182)" x2="0" y2="-14" stroke="#06982d"/><line aria-label="Date in 2009: 07/30; low: 24.85; high: 25.76" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(388.5000000348886,166.36363636363635)" x2="0" y2="-16.545454545454533" stroke="#ae1325"/><line aria-label="Date in 2009: 07/31; low: 24.93; high: 26.22" role="graphics-symbol" aria-roledescription="rule mark" transform="translate(395.00000003479454,164.9090909090909)" x2="0" y2="-23.45454545454544" stroke="#06982d"/></g><g class="mark-rect role-mark layer_1_marks" role="graphics-object" aria-roledescription="rect mark container"><path aria-label="Date in 2009: 06/01; open: 28.7; close: 30.04" role="graphics-symbol" aria-roledescription="bar" d="M2.5000000404369214,72.00000000000001h5v24.363636363636374h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 06/02; open: 30.04; close: 29.63" role="graphics-symbol" aria-roledescription="bar" d="M9.000000040342883,72.00000000000001h5v7.454545454545453h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 06/03; open: 29.62; close: 31.02" role="graphics-symbol" aria-roledescription="bar" d="M15.500000040248842,54.1818181818182h5v25.45454545454541h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 06/04; open: 31.02; close: 30.18" role="graphics-symbol" aria-roledescription="bar" d="M22.000000040154806,54.1818181818182h5v15.272727272727252h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 06/05; open: 29.39; close: 29.62" role="graphics-symbol" aria-roledescription="bar" d="M28.500000040060765,79.63636363636361h5v4.181818181818201h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 06/08; open: 30.84; close: 29.77" role="graphics-symbol" aria-roledescription="bar" d="M48.00000003977865,57.454545454545446h5v19.454545454545475h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 06/09; open: 29.77; close: 28.27" role="graphics-symbol" aria-roledescription="bar" d="M54.50000003968461,76.90909090909092h5v27.27272727272728h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 06/10; open: 26.9; close: 28.46" role="graphics-symbol" aria-roledescription="bar" d="M61.00000003959057,100.7272727272727h5v28.363636363636388h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 06/11; open: 27.36; close: 28.11" role="graphics-symbol" aria-roledescription="bar" d="M67.50000003949653,107.09090909090911h5v13.63636363636364h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 06/12; open: 28.08; close: 28.15" role="graphics-symbol" aria-roledescription="bar" d="M74.0000000394025,106.36363636363639h5v1.2727272727272805h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 06/15; open: 29.7; close: 30.81" role="graphics-symbol" aria-roledescription="bar" d="M93.50000003912037,58.00000000000003h5v20.181818181818173h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 06/16; open: 30.81; close: 32.68" role="graphics-symbol" aria-roledescription="bar" d="M100.00000003902633,24h5v34.00000000000003h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 06/17; open: 31.19; close: 31.54" role="graphics-symbol" aria-roledescription="bar" d="M106.50000003893228,44.72727272727275h5v6.363636363636331h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 06/18; open: 31.54; close: 30.03" role="graphics-symbol" aria-roledescription="bar" d="M113.00000003883825,44.72727272727275h5v27.454545454545425h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 06/19; open: 29.16; close: 27.99" role="graphics-symbol" aria-roledescription="bar" d="M119.5000000387442,87.99999999999999h5v21.27272727272731h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 06/22; open: 30.4; close: 31.17" role="graphics-symbol" aria-roledescription="bar" d="M139.0000000384621,51.45454545454542h5v14.000000000000064h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 06/23; open: 31.3; close: 30.58" role="graphics-symbol" aria-roledescription="bar" d="M145.50000003836806,49.09090909090907h5v13.090909090909129h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 06/24; open: 30.58; close: 29.05" role="graphics-symbol" aria-roledescription="bar" d="M152.000000038274,62.1818181818182h5v27.818181818181785h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 06/25; open: 29.45; close: 26.36" role="graphics-symbol" aria-roledescription="bar" d="M158.50000003818,82.72727272727273h5v56.18181818181817h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 06/26; open: 27.09; close: 25.93" role="graphics-symbol" aria-roledescription="bar" d="M165.00000003808594,125.63636363636364h5v21.09090909090908h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 06/29; open: 25.93; close: 25.35" role="graphics-symbol" aria-roledescription="bar" d="M184.50000003780383,146.72727272727272h5v10.545454545454533h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 06/30; open: 25.36; close: 26.35" role="graphics-symbol" aria-roledescription="bar" d="M191.00000003770978,139.09090909090907h5v18.00000000000003h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 07/01; open: 25.73; close: 26.22" role="graphics-symbol" aria-roledescription="bar" d="M197.50000003761573,141.45454545454547h5v8.909090909090878h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 07/02; open: 26.22; close: 27.95" role="graphics-symbol" aria-roledescription="bar" d="M204.00000003752172,110.00000000000001h5v31.454545454545453h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 07/06; open: 30.32; close: 29" role="graphics-symbol" aria-roledescription="bar" d="M230.00000003714555,66.90909090909089h5v24.00000000000003h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/07; open: 29; close: 30.85" role="graphics-symbol" aria-roledescription="bar" d="M236.50000003705148,57.27272727272725h5v33.63636363636367h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 07/08; open: 30.85; close: 31.3" role="graphics-symbol" aria-roledescription="bar" d="M243.00000003695746,49.09090909090907h5v8.18181818181818h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 07/09; open: 30.23; close: 29.78" role="graphics-symbol" aria-roledescription="bar" d="M249.50000003686344,68.54545454545453h5v8.181818181818173h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/10; open: 29.78; close: 29.02" role="graphics-symbol" aria-roledescription="bar" d="M256.0000000367694,76.7272727272727h5v13.818181818181841h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/13; open: 28.36; close: 26.31" role="graphics-symbol" aria-roledescription="bar" d="M275.5000000364873,102.54545454545453h5v37.27272727272731h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/14; open: 26.31; close: 25.02" role="graphics-symbol" aria-roledescription="bar" d="M282.0000000363932,139.81818181818184h5v23.45454545454544h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/15; open: 25.05; close: 25.89" role="graphics-symbol" aria-roledescription="bar" d="M288.5000000362992,147.45454545454547h5v15.272727272727252h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 07/16; open: 25.96; close: 25.42" role="graphics-symbol" aria-roledescription="bar" d="M295.00000003620517,146.18181818181816h5v9.818181818181813h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/17; open: 25.42; close: 24.34" role="graphics-symbol" aria-roledescription="bar" d="M301.50000003611115,155.99999999999997h5v19.636363636363654h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/20; open: 25.06; close: 24.4" role="graphics-symbol" aria-roledescription="bar" d="M321.000000035829,162.54545454545456h5v12h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/21; open: 24.28; close: 23.87" role="graphics-symbol" aria-roledescription="bar" d="M327.50000003573496,176.72727272727272h5v7.454545454545439h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/22; open: 24.05; close: 23.47" role="graphics-symbol" aria-roledescription="bar" d="M334.00000003564094,180.9090909090909h5v10.545454545454561h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/23; open: 23.71; close: 23.43" role="graphics-symbol" aria-roledescription="bar" d="M340.50000003554686,187.09090909090907h5v5.090909090909122h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/24; open: 23.87; close: 23.09" role="graphics-symbol" aria-roledescription="bar" d="M347.00000003545284,184.18181818181816h5v14.181818181818215h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/27; open: 24.06; close: 24.28" role="graphics-symbol" aria-roledescription="bar" d="M366.50000003517073,176.72727272727272h5v4.000000000000028h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 07/28; open: 24.28; close: 25.01" role="graphics-symbol" aria-roledescription="bar" d="M373.0000000350767,163.45454545454544h5v13.27272727272728h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 07/29; open: 25.47; close: 25.61" role="graphics-symbol" aria-roledescription="bar" d="M379.50000003498263,152.54545454545456h5v2.545454545454561h-5Z" fill="#06982d"/><path aria-label="Date in 2009: 07/30; open: 25.4; close: 25.4" role="graphics-symbol" aria-roledescription="bar" d="M386.0000000348886,156.3636363636364h5v0h-5Z" fill="#ae1325"/><path aria-label="Date in 2009: 07/31; open: 25.4; close: 25.92" role="graphics-symbol" aria-roledescription="bar" d="M392.50000003479454,146.90909090909088h5v9.454545454545524h-5Z" fill="#06982d"/></g></g><path class="foreground" aria-hidden="true" d="" display="none"/></g></g></g></svg>
}

func Example_compile() {
	vega := vegagoja.New()
	res, err := vega.Compile(candlestickSpec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	// Output:
	// {"$schema":"https://vega.github.io/schema/vega/v5.json","axes":[{"aria":false,"domain":false,"grid":true,"gridScale":"y","labels":false,"maxExtent":0,"minExtent":0,"orient":"bottom","scale":"x","tickCount":{"signal":"ceil(width/40)"},"ticks":false,"zindex":0},{"aria":false,"domain":false,"grid":true,"gridScale":"x","labels":false,"maxExtent":0,"minExtent":0,"orient":"left","scale":"y","tickCount":{"signal":"ceil(height/40)"},"ticks":false,"zindex":0},{"format":"%m/%d","grid":false,"labelAlign":"right","labelAngle":315,"labelBaseline":"top","labelFlush":true,"labelOverlap":true,"orient":"bottom","scale":"x","tickCount":{"signal":"ceil(width/40)"},"title":"Date in 2009","zindex":0},{"grid":false,"labelOverlap":true,"orient":"left","scale":"y","tickCount":{"signal":"ceil(height/40)"},"title":"Price","zindex":0}],"background":"white","data":[{"format":{"parse":{"date":"date"},"type":"json"},"name":"source_0","url":"data/ohlc.json"},{"name":"data_0","source":"source_0","transform":[{"expr":"(isDate(datum[\"date\"]) || (isValid(datum[\"date\"]) && isFinite(+datum[\"date\"]))) && isValid(datum[\"low\"]) && isFinite(+datum[\"low\"])","type":"filter"}]},{"name":"data_1","source":"source_0","transform":[{"expr":"(isDate(datum[\"date\"]) || (isValid(datum[\"date\"]) && isFinite(+datum[\"date\"]))) && isValid(datum[\"open\"]) && isFinite(+datum[\"open\"])","type":"filter"}]}],"description":"A candlestick chart inspired by an example in Protovis (http://mbostock.github.io/protovis/ex/candlestick.html)","height":200,"marks":[{"encode":{"update":{"description":{"signal":"\"Date in 2009: \" + (timeFormat(datum[\"date\"], '%m/%d')) + \"; low: \" + (format(datum[\"low\"], \"\")) + \"; high: \" + (format(datum[\"high\"], \"\"))"},"stroke":[{"test":"datum.open < datum.close","value":"#06982d"},{"value":"#ae1325"}],"x":{"field":"date","scale":"x"},"y":{"field":"low","scale":"y"},"y2":{"field":"high","scale":"y"}}},"from":{"data":"data_0"},"name":"layer_0_marks","style":["rule"],"type":"rule"},{"encode":{"update":{"ariaRoleDescription":{"value":"bar"},"description":{"signal":"\"Date in 2009: \" + (timeFormat(datum[\"date\"], '%m/%d')) + \"; open: \" + (format(datum[\"open\"], \"\")) + \"; close: \" + (format(datum[\"close\"], \"\"))"},"fill":[{"test":"datum.open < datum.close","value":"#06982d"},{"value":"#ae1325"}],"width":{"value":5},"xc":{"field":"date","scale":"x"},"y":{"field":"open","scale":"y"},"y2":{"field":"close","scale":"y"}}},"from":{"data":"data_1"},"name":"layer_1_marks","style":["bar"],"type":"rect"}],"padding":5,"scales":[{"domain":{"fields":[{"data":"data_0","field":"date"},{"data":"data_1","field":"date"}]},"name":"x","padding":5,"range":[0,{"signal":"width"}],"type":"time"},{"domain":{"fields":[{"data":"data_0","field":"low"},{"data":"data_0","field":"high"},{"data":"data_1","field":"open"},{"data":"data_1","field":"close"}]},"name":"y","nice":true,"range":[{"signal":"height"},0],"type":"linear","zero":false}],"style":"cell","width":400}
}

func Example_withParams() {
	vega := vegagoja.New(
		vegagoja.WithParams(map[string]interface{}{
			"x":            1,
			"x2":           100,
			"y":            1,
			"y2":           100,
			"cornerRadius": 20,
		}),
	)
	svg, err := vega.Render(context.Background(), squareSpec)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("square.svg", []byte(svg), 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(svg))
	// Output:
	// <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" class="marks" width="242" height="232" viewBox="0 0 242 232"><rect width="242" height="232" fill="white"/><g fill="none" stroke-miterlimit="10" transform="translate(36,10)"><g class="mark-group role-frame root" role="graphics-object" aria-roledescription="group mark container"><g transform="translate(0,0)"><path class="background" aria-hidden="true" d="M0.5,0.5h200v200h-200Z" stroke="#ddd"/><g><g class="mark-group role-axis" role="graphics-symbol" aria-roledescription="axis" aria-label="X-axis for a linear scale with values from 0 to 100"><g transform="translate(0.5,200.5)"><path class="background" aria-hidden="true" d="M0,0h0v0h0Z" pointer-events="none"/><g><g class="mark-rule role-axis-tick" pointer-events="none"><line transform="translate(0,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(40,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(80,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(120,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(160,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(200,0)" x2="0" y2="5" stroke="#888" stroke-width="1" opacity="1"/></g><g class="mark-text role-axis-label" pointer-events="none"><text text-anchor="start" transform="translate(0,15)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">0</text><text text-anchor="middle" transform="translate(40,15)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">20</text><text text-anchor="middle" transform="translate(80,15)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">40</text><text text-anchor="middle" transform="translate(120,15)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">60</text><text text-anchor="middle" transform="translate(160,15)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">80</text><text text-anchor="end" transform="translate(200,15)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">100</text></g><g class="mark-rule role-axis-domain" pointer-events="none"><line transform="translate(0,0)" x2="200" y2="0" stroke="#888" stroke-width="1" opacity="1"/></g></g><path class="foreground" aria-hidden="true" d="" pointer-events="none" display="none"/></g></g><g class="mark-group role-axis" role="graphics-symbol" aria-roledescription="axis" aria-label="Y-axis for a linear scale with values from 0 to 100"><g transform="translate(0.5,0.5)"><path class="background" aria-hidden="true" d="M0,0h0v0h0Z" pointer-events="none"/><g><g class="mark-rule role-axis-tick" pointer-events="none"><line transform="translate(0,200)" x2="-5" y2="0" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(0,160)" x2="-5" y2="0" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(0,120)" x2="-5" y2="0" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(0,80)" x2="-5" y2="0" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(0,40)" x2="-5" y2="0" stroke="#888" stroke-width="1" opacity="1"/><line transform="translate(0,0)" x2="-5" y2="0" stroke="#888" stroke-width="1" opacity="1"/></g><g class="mark-text role-axis-label" pointer-events="none"><text text-anchor="end" transform="translate(-7,203)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">0</text><text text-anchor="end" transform="translate(-7,163)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">20</text><text text-anchor="end" transform="translate(-7,123)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">40</text><text text-anchor="end" transform="translate(-7,83)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">60</text><text text-anchor="end" transform="translate(-7,42.99999999999999)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">80</text><text text-anchor="end" transform="translate(-7,3)" font-family="sans-serif" font-size="10px" fill="#000" opacity="1">100</text></g><g class="mark-rule role-axis-domain" pointer-events="none"><line transform="translate(0,200)" x2="0" y2="-200" stroke="#888" stroke-width="1" opacity="1"/></g></g><path class="foreground" aria-hidden="true" d="" pointer-events="none" display="none"/></g></g><g class="mark-rect role-mark marks" role="graphics-symbol" aria-roledescription="rect mark container"><path d="M22,0L180,0C191.03830048988,0,200,8.96169951012,200,20L200,178C200,189.03830048988,191.03830048988,198,180,198L22,198C10.96169951012,198,2,189.03830048988,2,178L2,20C2,8.96169951012,10.96169951012,0,22,0Z" fill="orange"/></g></g><path class="foreground" aria-hidden="true" d="" display="none"/></g></g></g></svg>
}

const squareSpec = `{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "params": [
    { "name": "x", "value": 25,
      "bind": {"input": "range", "min": 1, "max": 100, "step": 1} },
    { "name": "x2", "value": 75,
      "bind": {"input": "range", "min": 1, "max": 100, "step": 1} },
    { "name": "y", "value": 25,
      "bind": {"input": "range", "min": 1, "max": 100, "step": 1} },
    { "name": "y2", "value": 75,
      "bind": {"input": "range", "min": 1, "max": 100, "step": 1} },
    { "name": "cornerRadius", "value": 0,
      "bind": {"input": "range", "min": 0, "max": 50, "step": 1} }
  ],
  "data": {"values": [{}]},
  "mark": {
    "type": "rect",
    "color": "orange",
    "cornerRadius": {"expr": "cornerRadius"}
  },
  "encoding": {
    "x": {"datum": {"expr": "x"}, "type": "quantitative", "scale": {"domain": [0,100]}},
    "y": {"datum": {"expr": "y"}, "type": "quantitative", "scale": {"domain": [0,100]}},
    "x2": {"datum": {"expr": "x2"}},
    "y2": {"datum": {"expr": "y2"}}
  }
}`

func Example_withCSV() {
	vega := vegagoja.New(
		vegagoja.WithCSVString(co2Data),
	)
	svg, err := vega.Render(context.Background(), layerLineSpec)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("layer-line.svg", []byte(svg), 0o644); err != nil {
		log.Fatal(err)
	}
	// Output:
}

// candlestickSpec is the same as testdata/lite/layer_candlestick.vl.json
const candlestickSpec = `{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "width": 400,
  "description": "A candlestick chart inspired by an example in Protovis (http://mbostock.github.io/protovis/ex/candlestick.html)",
  "data": {"url": "data/ohlc.json"},
  "encoding": {
    "x": {
      "field": "date",
      "type": "temporal",
      "title": "Date in 2009",
      "axis": {
        "format": "%m/%d",
        "labelAngle": -45,
        "title": "Date in 2009"
      }
    },
    "y": {
      "type": "quantitative",
      "scale": {"zero": false},
      "axis": {"title": "Price"}
    },
    "color": {
      "condition": {
        "test": "datum.open < datum.close",
        "value": "#06982d"
      },
      "value": "#ae1325"
    }
  },
  "layer": [
    {
      "mark": "rule",
      "encoding": {
        "y": {"field": "low"},
        "y2": {"field": "high"}
      }
    },
    {
      "mark": "bar",
      "encoding": {
        "y": {"field": "open"},
        "y2": {"field": "close"}
      }
    }
  ]
}`

// layerLineSpec is same as testdata/lite/layer_line_co2_concentration.vl.json
const layerLineSpec = `{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "data": {
    "url": "data/co2-concentration.csv",
    "format": {"parse": {"Date": "utc:'%Y-%m-%d'"}}
  },
  "width": 800,
  "height": 500,
  "transform": [
    {"calculate": "year(datum.Date)", "as": "year"},
    {"calculate": "floor(datum.year / 10)", "as": "decade"},
    {
      "calculate": "(datum.year % 10) + (month(datum.Date)/12)",
      "as": "scaled_date"
    },
    {
      "calculate": "datum.first_date === datum.scaled_date ? 'first' : datum.last_date === datum.scaled_date ? 'last' : null",
      "as": "end"
    }
  ],
  "encoding": {
    "x": {
      "type": "quantitative",
      "title": "Year into Decade",
      "axis": {"tickCount": 11}
    },
    "y": {
      "title": "CO2 concentration in ppm",
      "type": "quantitative",
      "scale": {"zero": false}
    },
    "color": {
      "field": "decade",
      "type": "ordinal",
      "scale": {"scheme": "magma"},
      "legend": null
    }
  },

  "layer": [
    {
      "mark": "line",
      "encoding": {
        "x": {"field": "scaled_date"},
        "y": {"field": "CO2"}
      }
    },
    {
      "mark": {"type": "text", "baseline": "top", "aria": false},
      "encoding": {
        "x": {"aggregate": "min", "field": "scaled_date"},
        "y": {"aggregate": {"argmin": "scaled_date"}, "field": "CO2"},
        "text": {"aggregate": {"argmin": "scaled_date"}, "field": "year"}
      }
    },
    {
      "mark": {"type": "text", "aria": false},
      "encoding": {
        "x": {"aggregate": "max", "field": "scaled_date"},
        "y": {"aggregate": {"argmax": "scaled_date"}, "field": "CO2"},
        "text": {"aggregate": {"argmax": "scaled_date"}, "field": "year"}
      }
    }
  ],
  "config": {"text": {"align": "left", "dx": 3, "dy": 1}}
}`

const co2Data = `Date,CO2,adjusted CO2
1958-03-01,315.70,314.44
1958-04-01,317.46,315.16
1958-05-01,317.51,314.71
1958-07-01,315.86,315.19
1958-08-01,314.93,316.19
1958-09-01,313.21,316.08
1958-11-01,313.33,315.20
1958-12-01,314.67,315.43
1959-01-01,315.58,315.54
1959-02-01,316.49,315.86
1959-03-01,316.65,315.38
1959-04-01,317.72,315.42
1959-05-01,318.29,315.49
1959-06-01,318.15,316.03
1959-07-01,316.54,315.86
1959-08-01,314.80,316.06
1959-09-01,313.84,316.73
1959-10-01,313.33,316.33
1959-11-01,314.81,316.68
1959-12-01,315.58,316.35
1960-01-01,316.43,316.39
1960-02-01,316.98,316.35
1960-03-01,317.58,316.28
1960-04-01,319.03,316.70
1960-05-01,320.04,317.22
1960-06-01,319.58,317.48
1960-07-01,318.18,317.52
1960-08-01,315.90,317.20
1960-09-01,314.17,317.08
1960-10-01,313.83,316.83
1960-11-01,315.00,316.88
1960-12-01,316.19,316.96
1961-01-01,316.90,316.85
1961-02-01,317.70,317.07
1961-03-01,318.54,317.26
1961-04-01,319.48,317.16
1961-05-01,320.58,317.76
1961-06-01,319.77,317.63
1961-07-01,318.57,317.88
1961-08-01,316.79,318.06
1961-09-01,314.99,317.90
1961-10-01,315.31,318.32
1961-11-01,316.10,317.99
1961-12-01,317.01,317.78
1962-01-01,317.94,317.90
1962-02-01,318.55,317.92
1962-03-01,319.68,318.40
1962-04-01,320.57,318.24
1962-05-01,321.02,318.18
1962-06-01,320.62,318.47
1962-07-01,319.61,318.93
1962-08-01,317.40,318.68
1962-09-01,316.25,319.17
1962-10-01,315.42,318.44
1962-11-01,316.69,318.58
1962-12-01,317.70,318.47
1963-01-01,318.74,318.70
1963-02-01,319.07,318.43
1963-03-01,319.86,318.57
1963-04-01,321.38,319.05
1963-05-01,322.25,319.40
1963-06-01,321.48,319.33
1963-07-01,319.74,319.06
1963-08-01,317.77,319.05
1963-09-01,316.21,319.14
1963-10-01,315.99,319.02
1963-11-01,317.07,318.97
1963-12-01,318.35,319.13
1964-01-01,319.57,319.53
1964-05-01,322.26,319.40
1964-06-01,321.89,319.75
1964-07-01,320.44,319.78
1964-08-01,318.69,320.01
1964-09-01,316.70,319.66
1964-10-01,316.87,319.91
1964-11-01,317.68,319.58
1964-12-01,318.71,319.49
1965-01-01,319.44,319.40
1965-02-01,320.44,319.81
1965-03-01,320.89,319.59
1965-04-01,322.14,319.78
1965-05-01,322.17,319.30
1965-06-01,321.87,319.70
1965-07-01,321.21,320.52
1965-08-01,318.87,320.16
1965-09-01,317.81,320.77
1965-10-01,317.30,320.36
1965-11-01,318.87,320.78
1965-12-01,319.42,320.20
1966-01-01,320.62,320.58
1966-02-01,321.60,320.95
1966-03-01,322.39,321.08
1966-04-01,323.70,321.34
1966-05-01,324.08,321.20
1966-06-01,323.75,321.57
1966-07-01,322.38,321.69
1966-08-01,320.36,321.66
1966-09-01,318.64,321.60
1966-10-01,318.10,321.17
1966-11-01,319.78,321.70
1966-12-01,321.03,321.81
1967-01-01,322.33,322.29
1967-02-01,322.50,321.85
1967-03-01,323.04,321.73
1967-04-01,324.42,322.05
1967-05-01,325.00,322.11
1967-06-01,324.09,321.91
1967-07-01,322.54,321.85
1967-08-01,320.92,322.22
1967-09-01,319.25,322.23
1967-10-01,319.39,322.47
1967-11-01,320.73,322.65
1967-12-01,321.96,322.74
1968-01-01,322.57,322.53
1968-02-01,323.15,322.50
1968-03-01,323.89,322.55
1968-04-01,325.02,322.62
1968-05-01,325.57,322.68
1968-06-01,325.36,323.19
1968-07-01,324.14,323.47
1968-08-01,322.11,323.44
1968-09-01,320.33,323.32
1968-10-01,320.25,323.33
1968-11-01,321.32,323.25
1968-12-01,322.89,323.68
1969-01-01,324.00,323.96
1969-02-01,324.41,323.77
1969-03-01,325.63,324.32
1969-04-01,326.66,324.28
1969-05-01,327.38,324.48
1969-06-01,326.71,324.51
1969-07-01,325.88,325.18
1969-08-01,323.66,324.97
1969-09-01,322.38,325.37
1969-10-01,321.78,324.88
1969-11-01,322.85,324.79
1969-12-01,324.12,324.91
1970-01-01,325.06,325.02
1970-02-01,325.98,325.33
1970-03-01,326.93,325.61
1970-04-01,328.13,325.74
1970-05-01,328.08,325.16
1970-06-01,327.67,325.46
1970-07-01,326.34,325.64
1970-08-01,324.69,326.00
1970-09-01,323.10,326.10
1970-10-01,323.06,326.18
1970-11-01,324.01,325.95
1970-12-01,325.13,325.93
1971-01-01,326.17,326.13
1971-02-01,326.68,326.03
1971-03-01,327.17,325.85
1971-04-01,327.79,325.38
1971-05-01,328.93,326.00
1971-06-01,328.57,326.36
1971-07-01,327.36,326.65
1971-08-01,325.43,326.75
1971-09-01,323.36,326.37
1971-10-01,323.56,326.68
1971-11-01,324.80,326.75
1971-12-01,326.01,326.81
1972-01-01,326.77,326.73
1972-02-01,327.63,326.98
1972-03-01,327.75,326.40
1972-04-01,329.72,327.29
1972-05-01,330.07,327.13
1972-06-01,329.09,326.89
1972-07-01,328.04,327.36
1972-08-01,326.32,327.67
1972-09-01,324.84,327.87
1972-10-01,325.20,328.33
1972-11-01,326.50,328.45
1972-12-01,327.55,328.35
1973-01-01,328.55,328.50
1973-02-01,329.56,328.90
1973-03-01,330.30,328.97
1973-04-01,331.50,329.08
1973-05-01,332.48,329.53
1973-06-01,332.07,329.84
1973-07-01,330.87,330.16
1973-08-01,329.31,330.64
1973-09-01,327.51,330.55
1973-10-01,327.18,330.32
1973-11-01,328.16,330.13
1973-12-01,328.64,329.44
1974-01-01,329.35,329.31
1974-02-01,330.71,330.05
1974-03-01,331.48,330.14
1974-04-01,332.65,330.22
1974-05-01,333.09,330.13
1974-06-01,332.25,330.01
1974-07-01,331.18,330.46
1974-08-01,329.39,330.73
1974-09-01,327.43,330.48
1974-10-01,327.37,330.52
1974-11-01,328.46,330.43
1974-12-01,329.57,330.38
1975-01-01,330.40,330.36
1975-02-01,331.40,330.74
1975-03-01,332.04,330.69
1975-04-01,333.31,330.87
1975-05-01,333.97,331.00
1975-06-01,333.60,331.36
1975-07-01,331.90,331.19
1975-08-01,330.06,331.39
1975-09-01,328.56,331.61
1975-10-01,328.34,331.50
1975-11-01,329.49,331.47
1975-12-01,330.76,331.57
1976-01-01,331.75,331.70
1976-02-01,332.57,331.90
1976-03-01,333.50,332.12
1976-04-01,334.58,332.12
1976-05-01,334.88,331.90
1976-06-01,334.33,332.10
1976-07-01,333.05,332.36
1976-08-01,330.94,332.31
1976-09-01,329.30,332.38
1976-10-01,328.94,332.11
1976-11-01,330.31,332.29
1976-12-01,331.68,332.49
1977-01-01,332.93,332.88
1977-02-01,333.42,332.75
1977-03-01,334.70,333.35
1977-04-01,336.07,333.62
1977-05-01,336.75,333.76
1977-06-01,336.27,334.01
1977-07-01,334.92,334.20
1977-08-01,332.75,334.10
1977-09-01,331.59,334.67
1977-10-01,331.16,334.35
1977-11-01,332.40,334.40
1977-12-01,333.85,334.66
1978-01-01,334.97,334.93
1978-02-01,335.39,334.72
1978-03-01,336.64,335.28
1978-04-01,337.76,335.30
1978-05-01,338.01,335.02
1978-06-01,337.90,335.63
1978-07-01,336.54,335.82
1978-08-01,334.68,336.03
1978-09-01,332.76,335.85
1978-10-01,332.55,335.74
1978-11-01,333.92,335.92
1978-12-01,334.95,335.77
1979-01-01,336.23,336.18
1979-02-01,336.76,336.09
1979-03-01,337.96,336.60
1979-04-01,338.88,336.41
1979-05-01,339.47,336.47
1979-06-01,339.29,337.01
1979-07-01,337.73,337.01
1979-08-01,336.09,337.44
1979-09-01,333.92,337.01
1979-10-01,333.86,337.07
1979-11-01,335.29,337.30
1979-12-01,336.73,337.55
1980-01-01,338.01,337.97
1980-02-01,338.36,337.69
1980-03-01,340.07,338.68
1980-04-01,340.76,338.26
1980-05-01,341.47,338.45
1980-06-01,341.17,338.91
1980-07-01,339.56,338.86
1980-08-01,337.60,338.99
1980-09-01,335.88,339.00
1980-10-01,336.02,339.23
1980-11-01,337.10,339.11
1980-12-01,338.21,339.03
1981-01-01,339.24,339.19
1981-02-01,340.48,339.80
1981-03-01,341.38,340.01
1981-04-01,342.50,340.02
1981-05-01,342.91,339.89
1981-06-01,342.25,339.96
1981-07-01,340.49,339.76
1981-08-01,338.43,339.80
1981-09-01,336.69,339.81
1981-10-01,336.86,340.08
1981-11-01,338.36,340.38
1981-12-01,339.61,340.44
1982-01-01,340.75,340.71
1982-02-01,341.61,340.94
1982-03-01,342.70,341.32
1982-04-01,343.57,341.08
1982-05-01,344.14,341.10
1982-06-01,343.35,341.05
1982-07-01,342.06,341.32
1982-08-01,339.81,341.18
1982-09-01,337.98,341.10
1982-10-01,337.86,341.10
1982-11-01,339.26,341.29
1982-12-01,340.49,341.32
1983-01-01,341.38,341.33
1983-02-01,342.52,341.84
1983-03-01,343.10,341.72
1983-04-01,344.94,342.44
1983-05-01,345.76,342.71
1983-06-01,345.32,343.01
1983-07-01,343.98,343.25
1983-08-01,342.38,343.75
1983-09-01,339.87,343.00
1983-10-01,339.99,343.24
1983-11-01,341.15,343.19
1983-12-01,342.99,343.82
1984-01-01,343.70,343.65
1984-02-01,344.50,343.83
1984-03-01,345.28,343.87
1984-04-01,347.05,344.52
1984-05-01,347.43,344.38
1984-06-01,346.80,344.51
1984-07-01,345.39,344.69
1984-08-01,343.28,344.68
1984-09-01,341.07,344.23
1984-10-01,341.35,344.60
1984-11-01,342.98,345.01
1984-12-01,344.22,345.05
1985-01-01,344.97,344.92
1985-02-01,345.99,345.31
1985-03-01,347.42,346.04
1985-04-01,348.35,345.83
1985-05-01,348.93,345.86
1985-06-01,348.25,345.93
1985-07-01,346.56,345.82
1985-08-01,344.67,346.06
1985-09-01,343.09,346.24
1985-10-01,342.80,346.07
1985-11-01,344.24,346.29
1985-12-01,345.56,346.39
1986-01-01,346.30,346.25
1986-02-01,346.95,346.27
1986-03-01,347.85,346.46
1986-04-01,349.55,347.03
1986-05-01,350.22,347.14
1986-06-01,349.55,347.23
1986-07-01,347.94,347.20
1986-08-01,345.90,347.29
1986-09-01,344.85,348.02
1986-10-01,344.17,347.45
1986-11-01,345.66,347.71
1986-12-01,346.90,347.74
1987-01-01,348.02,347.98
1987-02-01,348.48,347.79
1987-03-01,349.42,348.02
1987-04-01,350.98,348.45
1987-05-01,351.85,348.76
1987-06-01,351.26,348.92
1987-07-01,349.51,348.77
1987-08-01,348.10,349.49
1987-09-01,346.45,349.62
1987-10-01,346.36,349.65
1987-11-01,347.81,349.87
1987-12-01,348.96,349.81
1988-01-01,350.43,350.39
1988-02-01,351.73,351.04
1988-03-01,352.22,350.79
1988-04-01,353.59,351.02
1988-05-01,354.22,351.12
1988-06-01,353.80,351.48
1988-07-01,352.38,351.66
1988-08-01,350.43,351.85
1988-09-01,348.73,351.92
1988-10-01,348.88,352.18
1988-11-01,350.07,352.13
1988-12-01,351.34,352.18
1989-01-01,352.76,352.71
1989-02-01,353.07,352.38
1989-03-01,353.68,352.27
1989-04-01,355.42,352.87
1989-05-01,355.67,352.56
1989-06-01,355.12,352.77
1989-07-01,353.90,353.15
1989-08-01,351.67,353.07
1989-09-01,349.81,353.00
1989-10-01,349.99,353.30
1989-11-01,351.30,353.37
1989-12-01,352.52,353.37
1990-01-01,353.66,353.62
1990-02-01,354.70,354.00
1990-03-01,355.38,353.97
1990-04-01,356.20,353.64
1990-05-01,357.16,354.04
1990-06-01,356.23,353.87
1990-07-01,354.81,354.06
1990-08-01,352.91,354.31
1990-09-01,350.96,354.17
1990-10-01,351.18,354.50
1990-11-01,352.83,354.91
1990-12-01,354.21,355.06
1991-01-01,354.72,354.68
1991-02-01,355.75,355.05
1991-03-01,357.16,355.74
1991-04-01,358.60,356.03
1991-05-01,359.34,356.21
1991-06-01,358.24,355.88
1991-07-01,356.17,355.42
1991-08-01,354.01,355.42
1991-09-01,352.15,355.37
1991-10-01,352.21,355.55
1991-11-01,353.75,355.83
1991-12-01,354.99,355.84
1992-01-01,355.99,355.94
1992-02-01,356.72,356.02
1992-03-01,357.81,356.36
1992-04-01,359.15,356.55
1992-05-01,359.66,356.53
1992-06-01,359.25,356.90
1992-07-01,357.02,356.30
1992-08-01,355.00,356.44
1992-09-01,353.01,356.25
1992-10-01,353.31,356.64
1992-11-01,354.16,356.25
1992-12-01,355.40,356.25
1993-01-01,356.70,356.66
1993-02-01,357.17,356.46
1993-03-01,358.38,356.95
1993-04-01,359.46,356.88
1993-05-01,360.28,357.13
1993-06-01,359.60,357.22
1993-07-01,357.57,356.81
1993-08-01,355.52,356.94
1993-09-01,353.69,356.93
1993-10-01,353.99,357.34
1993-11-01,355.34,357.44
1993-12-01,356.80,357.66
1994-01-01,358.37,358.32
1994-02-01,358.91,358.21
1994-03-01,359.97,358.54
1994-04-01,361.26,358.67
1994-05-01,361.69,358.53
1994-06-01,360.94,358.56
1994-07-01,359.55,358.79
1994-08-01,357.48,358.90
1994-09-01,355.84,359.09
1994-10-01,356.00,359.36
1994-11-01,357.58,359.69
1994-12-01,359.04,359.90
1995-01-01,359.97,359.92
1995-02-01,361.00,360.30
1995-03-01,361.63,360.20
1995-04-01,363.45,360.85
1995-05-01,363.80,360.63
1995-06-01,363.26,360.87
1995-07-01,361.89,361.13
1995-08-01,359.45,360.88
1995-09-01,358.05,361.31
1995-10-01,357.75,361.13
1995-11-01,359.56,361.68
1995-12-01,360.70,361.56
1996-01-01,362.05,362.00
1996-02-01,363.24,362.54
1996-03-01,364.02,362.56
1996-04-01,364.71,362.08
1996-05-01,365.42,362.24
1996-06-01,364.97,362.59
1996-07-01,363.65,362.92
1996-08-01,361.48,362.94
1996-09-01,359.45,362.73
1996-10-01,359.61,362.99
1996-11-01,360.76,362.87
1996-12-01,362.33,363.19
1997-01-01,363.19,363.14
1997-02-01,363.99,363.28
1997-03-01,364.56,363.12
1997-04-01,366.36,363.74
1997-05-01,366.80,363.61
1997-06-01,365.63,363.22
1997-07-01,364.47,363.70
1997-08-01,362.50,363.94
1997-09-01,360.19,363.47
1997-10-01,360.78,364.17
1997-11-01,362.43,364.56
1997-12-01,364.28,365.14
1998-01-01,365.33,365.28
1998-02-01,366.15,365.44
1998-03-01,367.31,365.87
1998-04-01,368.61,365.99
1998-05-01,369.30,366.11
1998-06-01,368.88,366.46
1998-07-01,367.64,366.87
1998-08-01,365.78,367.22
1998-09-01,363.90,367.19
1998-10-01,364.23,367.64
1998-11-01,365.46,367.59
1998-12-01,366.97,367.84
1999-01-01,368.15,368.10
1999-02-01,368.87,368.16
1999-03-01,369.59,368.14
1999-04-01,371.14,368.51
1999-05-01,371.00,367.80
1999-06-01,370.35,367.93
1999-07-01,369.27,368.49
1999-08-01,366.93,368.37
1999-09-01,364.64,367.94
1999-10-01,365.13,368.55
1999-11-01,366.68,368.81
1999-12-01,368.00,368.88
2000-01-01,369.14,369.09
2000-02-01,369.46,368.75
2000-03-01,370.51,369.03
2000-04-01,371.66,369.00
2000-05-01,371.83,368.61
2000-06-01,371.69,369.28
2000-07-01,370.12,369.37
2000-08-01,368.12,369.60
2000-09-01,366.62,369.94
2000-10-01,366.73,370.15
2000-11-01,368.29,370.43
2000-12-01,369.52,370.40
2001-01-01,370.28,370.23
2001-02-01,371.50,370.78
2001-03-01,372.12,370.66
2001-04-01,372.86,370.21
2001-05-01,374.02,370.79
2001-06-01,373.31,370.87
2001-07-01,371.62,370.84
2001-08-01,369.55,371.00
2001-09-01,367.96,371.28
2001-10-01,368.09,371.53
2001-11-01,369.68,371.83
2001-12-01,371.24,372.12
2002-01-01,372.44,372.39
2002-02-01,373.08,372.36
2002-03-01,373.52,372.05
2002-04-01,374.85,372.20
2002-05-01,375.55,372.31
2002-06-01,375.40,372.95
2002-07-01,374.02,373.24
2002-08-01,371.48,372.94
2002-09-01,370.70,374.03
2002-10-01,370.25,373.70
2002-11-01,372.08,374.24
2002-12-01,373.78,374.66
2003-01-01,374.68,374.63
2003-02-01,375.62,374.90
2003-03-01,376.11,374.64
2003-04-01,377.65,374.99
2003-05-01,378.35,375.11
2003-06-01,378.13,375.67
2003-07-01,376.60,375.82
2003-08-01,374.48,375.95
2003-09-01,372.98,376.32
2003-10-01,373.00,376.46
2003-11-01,374.35,376.51
2003-12-01,375.69,376.57
2004-01-01,376.79,376.74
2004-02-01,377.37,376.64
2004-03-01,378.39,376.89
2004-04-01,380.50,377.80
2004-05-01,380.62,377.36
2004-06-01,379.55,377.11
2004-07-01,377.76,377.01
2004-08-01,375.83,377.33
2004-09-01,374.05,377.41
2004-10-01,374.22,377.69
2004-11-01,375.84,378.01
2004-12-01,377.44,378.33
2005-01-01,378.34,378.29
2005-02-01,379.61,378.88
2005-03-01,380.08,378.61
2005-04-01,382.05,379.37
2005-05-01,382.24,378.97
2005-06-01,382.08,379.61
2005-07-01,380.66,379.88
2005-08-01,378.67,380.14
2005-09-01,376.42,379.78
2005-10-01,376.80,380.28
2005-11-01,378.31,380.49
2005-12-01,379.96,380.85
2006-01-01,381.37,381.32
2006-02-01,382.02,381.29
2006-03-01,382.56,381.08
2006-04-01,384.36,381.68
2006-05-01,384.92,381.65
2006-06-01,384.03,381.55
2006-07-01,382.28,381.49
2006-08-01,380.48,381.95
2006-09-01,378.81,382.18
2006-10-01,379.06,382.55
2006-11-01,380.14,382.33
2006-12-01,381.66,382.55
2007-01-01,382.58,382.53
2007-02-01,383.71,382.98
2007-03-01,384.34,382.85
2007-04-01,386.23,383.53
2007-05-01,386.41,383.13
2007-06-01,385.87,383.39
2007-07-01,384.44,383.65
2007-08-01,381.84,383.32
2007-09-01,380.86,384.25
2007-10-01,380.86,384.36
2007-11-01,382.36,384.55
2007-12-01,383.61,384.51
2008-01-01,385.07,385.02
2008-02-01,385.84,385.11
2008-03-01,385.83,384.31
2008-04-01,386.77,384.04
2008-05-01,388.51,385.22
2008-06-01,388.05,385.58
2008-07-01,386.25,385.49
2008-08-01,384.08,385.60
2008-09-01,383.09,386.49
2008-10-01,382.78,386.28
2008-11-01,384.01,386.20
2008-12-01,385.11,386.01
2009-01-01,386.65,386.61
2009-02-01,387.12,386.39
2009-03-01,388.52,387.02
2009-04-01,389.57,386.86
2009-05-01,390.17,386.86
2009-06-01,389.62,387.12
2009-07-01,388.07,387.27
2009-08-01,386.08,387.57
2009-09-01,384.65,388.06
2009-10-01,384.33,387.85
2009-11-01,386.05,388.25
2009-12-01,387.49,388.39
2010-01-01,388.55,388.50
2010-02-01,390.08,389.34
2010-03-01,391.01,389.51
2010-04-01,392.38,389.66
2010-05-01,393.22,389.90
2010-06-01,392.24,389.73
2010-07-01,390.33,389.53
2010-08-01,388.52,390.01
2010-09-01,386.84,390.25
2010-10-01,387.16,390.70
2010-11-01,388.67,390.88
2010-12-01,389.81,390.71
2011-01-01,391.30,391.25
2011-02-01,391.92,391.18
2011-03-01,392.45,390.95
2011-04-01,393.37,390.64
2011-05-01,394.28,390.96
2011-06-01,393.69,391.18
2011-07-01,392.59,391.79
2011-08-01,390.21,391.71
2011-09-01,389.00,392.43
2011-10-01,388.93,392.48
2011-11-01,390.24,392.46
2011-12-01,391.80,392.71
2012-01-01,393.07,393.02
2012-02-01,393.35,392.61
2012-03-01,394.36,392.82
2012-04-01,396.43,393.66
2012-05-01,396.87,393.53
2012-06-01,395.88,393.38
2012-07-01,394.52,393.75
2012-08-01,392.54,394.07
2012-09-01,391.13,394.57
2012-10-01,391.01,394.56
2012-11-01,392.95,395.17
2012-12-01,394.33,395.24
2013-01-01,395.61,395.55
2013-02-01,396.85,396.10
2013-03-01,397.26,395.74
2013-04-01,398.35,395.60
2013-05-01,399.98,396.63
2013-06-01,398.87,396.34
2013-07-01,397.37,396.56
2013-08-01,395.41,396.92
2013-09-01,393.39,396.84
2013-10-01,393.70,397.26
2013-11-01,395.19,397.43
2013-12-01,396.82,397.73
2014-01-01,397.93,397.87
2014-02-01,398.10,397.35
2014-03-01,399.47,397.95
2014-04-01,401.33,398.57
2014-05-01,401.88,398.52
2014-06-01,401.31,398.77
2014-07-01,399.07,398.26
2014-08-01,397.21,398.72
2014-09-01,395.40,398.86
2014-10-01,395.65,399.23
2014-11-01,397.23,399.46
2014-12-01,398.79,399.70
2015-01-01,399.85,399.80
2015-02-01,400.31,399.56
2015-03-01,401.51,399.99
2015-04-01,403.45,400.69
2015-05-01,404.10,400.74
2015-06-01,402.88,400.33
2015-07-01,401.61,400.80
2015-08-01,399.00,400.51
2015-09-01,397.50,400.96
2015-10-01,398.28,401.87
2015-11-01,400.24,402.48
2015-12-01,401.89,402.81
2016-01-01,402.65,402.60
2016-02-01,404.16,403.41
2016-03-01,404.85,403.30
2016-04-01,407.57,404.77
2016-05-01,407.66,404.28
2016-06-01,407.00,404.48
2016-07-01,404.50,403.72
2016-08-01,402.24,403.79
2016-09-01,401.01,404.50
2016-10-01,401.50,405.09
2016-11-01,403.64,405.88
2016-12-01,404.55,405.47
2017-01-01,406.07,406.02
2017-02-01,406.64,405.89
2017-03-01,407.05,405.52
2017-04-01,408.95,406.17
2017-05-01,409.91,406.52
2017-06-01,409.12,406.56
2017-07-01,407.20,406.38
2017-08-01,405.24,406.76
2017-09-01,403.27,406.75
2017-10-01,403.64,407.25
2017-11-01,405.17,407.43
2017-12-01,406.75,407.68
2018-01-01,408.05,408.00
2018-02-01,408.34,407.59
2018-03-01,409.25,407.72
2018-04-01,410.30,407.52
2018-05-01,411.30,407.91
2018-06-01,410.88,408.31
2018-07-01,408.90,408.08
2018-08-01,407.10,408.63
2018-09-01,405.59,409.08
2018-10-01,405.99,409.61
2018-11-01,408.12,410.38
2018-12-01,409.23,410.15
2019-01-01,410.92,410.87
2019-02-01,411.66,410.90
2019-03-01,412.00,410.46
2019-04-01,413.52,410.72
2019-05-01,414.83,411.42
2019-06-01,413.96,411.38
2019-07-01,411.85,411.03
2019-08-01,410.08,411.62
2019-09-01,408.55,412.06
2019-10-01,408.43,412.06
2019-11-01,410.29,412.56
2019-12-01,411.85,412.78
2020-01-01,413.37,413.32
2020-02-01,414.09,413.33
2020-03-01,414.51,412.94
2020-04-01,416.18,413.35
`
