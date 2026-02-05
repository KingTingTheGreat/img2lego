package converter

type Color struct {
	Hex  string
	Name string
}

var BG_COLORS []Color = []Color{
	{Hex: "#05131D", Name: "Black"},
	{Hex: "#FFFFFF", Name: "White"},
	{Hex: "#F6D7B3", Name: "LightNougat"},
	{Hex: "#9B9A5A", Name: "OliveGreen"},
	{Hex: "#E1D5ED", Name: "Lavender"},
}

var LEGO_COLORS []Color = []Color{
	{Hex: "#4B9F4A", Name: "BrightGreen"},
	{Hex: "#36AEBF", Name: "MediumAzur"},          // medium azure
	{Hex: "#923978", Name: "BrightReddishViolet"}, // magenta
	{Hex: "#352100", Name: "DarkBrown"},
	{Hex: "#008F9B", Name: "BrightBluishGreen"}, // dark turquoise
	{Hex: "#0055BF", Name: "BrightBlue"},        // blue
	{Hex: "#AA7D55", Name: "MediumNougat"},
	{Hex: "#F2CD37", Name: "BrightYellow"}, // yellow
	{Hex: "#C91A09", Name: "BrightRed"},    // red
	{Hex: "#E4CD9E", Name: "BrickYellow"},  // tan
	{Hex: "#FFFFFF", Name: "White"},
	{Hex: "#05131D", Name: "Black"},
	{Hex: "#0A3463", Name: "EarthBlue"},  // dark blue
	{Hex: "#FFF03A", Name: "CoolYellow"}, // bright light yellow
	{Hex: "#5A93DB", Name: "MediumBlue"},
	{Hex: "#DFEEA5", Name: "SpringYellowishGreen"}, // yellowish green
	{Hex: "#E4ADC8", Name: "LightPurple"},          // bright pink
	{Hex: "#BBE90B", Name: "BrightYellowishGreen"}, // lime
	{Hex: "#6074A1", Name: "SandBlue"},
	{Hex: "#958A73", Name: "SandYellow"}, // dark tan
	{Hex: "#A95500", Name: "DarkOrange"},
	{Hex: "#078BC9", Name: "DarkAzur"},             // dark azure
	{Hex: "#9FC3E9", Name: "LightRoyalBlue"},       // bright light blue
	{Hex: "#F8BB3D", Name: "FlameYellowishOrange"}, // bright light orange
	{Hex: "#C870A0", Name: "BrightPurple"},         // dark pink
	{Hex: "#E1D5ED", Name: "Lavender"},
	{Hex: "#B3D7D1", Name: "Aqua"},
	{Hex: "#FF698F", Name: "VibrantCoral"},   // coral
	{Hex: "#898788", Name: "SilverMetallic"}, // flat silver
	{Hex: "#AA7F2E", Name: "WarmGold"},       // pearl gold
	{Hex: "#FE8A18", Name: "BrightOrange"},   // orange
	{Hex: "#F6D7B3", Name: "LightNougat"},
	{Hex: "#6C6E68", Name: "DarkStoneGrey"}, // dark bluish gray
	{Hex: "#582A12", Name: "ReddishBrown"},
	{Hex: "#720E0F", Name: "DarkRed"},
	{Hex: "#A0A5A9", Name: "MediumStoneGrey"}, // light bluish gray
	{Hex: "#D9D9D9", Name: "WhiteGlow"},       // glow in dark white
}

var BRICKLINK_COLORS []Color = []Color{
	{Hex: "#05131D", Name: "Black"},
	{Hex: "#0055BF", Name: "Blue"},
	{Hex: "#4B9F4A", Name: "BrightGreen"},
	{Hex: "#9FC3E9", Name: "BrightLightBlue"},
	{Hex: "#F8BB3D", Name: "BrightLightOrange"},
	{Hex: "#FFF03A", Name: "BrightLightYellow"},
	{Hex: "#E4ADC8", Name: "TransPink"},
	{Hex: "#583927", Name: "Brown"},
	{Hex: "#FF698F", Name: "Coral"},
	{Hex: "#078BC9", Name: "DarkAzure"},
	{Hex: "#0A3463", Name: "DarkBlue"},
	{Hex: "#6C6E68", Name: "DarkBluishGray"},
	{Hex: "#352100", Name: "DarkBrown"},
	{Hex: "#184632", Name: "DarkGreen"},
	{Hex: "#A95500", Name: "DarkOrange"},
	{Hex: "#C870A0", Name: "DarkPink"},
	{Hex: "#720E0F", Name: "DarkRed"},
	{Hex: "#958A73", Name: "DarkTan"},
	{Hex: "#008F9B", Name: "DarkTurquoise"},
	{Hex: "#237841", Name: "Green"},
	{Hex: "#E1D5ED", Name: "Lavender"},
	{Hex: "#ADC3C0", Name: "LightAqua"},
	{Hex: "#A0A5A9", Name: "LightBluishGray"},
	{Hex: "#9BA19D", Name: "LightGray"},
	{Hex: "#F6D7B3", Name: "LightNougat"},
	{Hex: "#FECCCF", Name: "LightPink"},
	{Hex: "#BBE90B", Name: "Lime"},
	{Hex: "#923978", Name: "Magenta"},
	{Hex: "#36AEBF", Name: "MediumAzure"},
	{Hex: "#5A93DB", Name: "MediumBlue"},
	{Hex: "#AC78BA", Name: "MediumLavender"},
	{Hex: "#AA7D55", Name: "MediumNougat"},
	{Hex: "#FFA70B", Name: "MediumOrange"},
	{Hex: "#EBD800", Name: "NeonYellow"}, // vibrant yellow
	{Hex: "#D09168", Name: "Nougat"},
	{Hex: "#9B9A5A", Name: "OliveGreen"},
	{Hex: "#FE8A18", Name: "Orange"},
	{Hex: "#FC97AC", Name: "Pink"},
	{Hex: "#C91A09", Name: "TransRed"},
	{Hex: "#582A12", Name: "ReddishBrown"},
	{Hex: "#6074A1", Name: "SandBlue"},
	{Hex: "#E4CD9E", Name: "Tan"},
	{Hex: "#FFFFFF", Name: "GlitterTransClear"},
	{Hex: "#F2CD37", Name: "Yellow"},
	{Hex: "#DFEEA5", Name: "YellowishGreen"},
	{Hex: "#635F52", Name: "TransBlack"},
	{Hex: "#D9E4A7", Name: "TransBrightGreen"},
	{Hex: "#FCFCFC", Name: "SatinTransClear"},    // trans clear opal
	{Hex: "#0020A0", Name: "SatinTransDarkBlue"}, // trans dark blue opal
	{Hex: "#DF6695", Name: "GlitterTransDarkPink"},
	{Hex: "#84B68D", Name: "SatinTransBrightGreen"},
	{Hex: "#AEEFEC", Name: "TransLightBlue"}, // trans blue opal
	{Hex: "#C9E788", Name: "TransLightBrightGreen"},
	{Hex: "#FCB76D", Name: "TransLightOrange"}, // trans flame yellowish orange
	{Hex: "#F8F184", Name: "TransNeonGreen"},
	{Hex: "#FF800D", Name: "TransNeonOrange"},
	{Hex: "#DAB000", Name: "TransNeonYellow"},
	{Hex: "#F08F1C", Name: "TransOrange"},
	{Hex: "#A5A5CB", Name: "TransPurple"},
	{Hex: "#F5CD2F", Name: "TransYellow"},
	{Hex: "#E0E0E0", Name: "ChromeSilver"},
	{Hex: "#B48455", Name: "FlatDarkGold"},
	{Hex: "#898788", Name: "FlatSilver"},
	{Hex: "#575857", Name: "PearlDarkGray"},
	{Hex: "#AA7F2E", Name: "PearlGold"},
	{Hex: "#DCBC81", Name: "PearlLightGold"},
	{Hex: "#68BCC5", Name: "GlitterTransLightBlue"},
	{Hex: "#DBAC34", Name: "MetallicGold"},
	{Hex: "#D4D5C9", Name: "GlowInDarkOpaque"},
	{Hex: "#D9D9D9", Name: "GlowInDarkWhite"},
}

var PART_NUMBERS []Color = []Color{
	{Hex: "#352100", Name: "6322813"},  // dark brown
	{Hex: "#6074A1", Name: "6322842"},  // sand blue
	{Hex: "#B3D7D1", Name: "6322818"},  // aqua
	{Hex: "#FE8A18", Name: "6284582"},  // bright orange, orange
	{Hex: "#FF698F", Name: "6311436"},  // coral, vibrant coral
	{Hex: "#F8BB3D", Name: "6322822"},  // flame yellowish orange, bright light orange
	{Hex: "#9B9A5A", Name: "6284595"},  // olive green
	{Hex: "#BBE90B", Name: "6284583"},  // bright yellowish green, lime
	{Hex: "#720E0F", Name: "6284585"},  // dark red
	{Hex: "#68BCC5", Name: "6299918"},  // glitter trans light blue
	{Hex: "#FFFFFF", Name: "6284572"},  // white, glitter trans clear
	{Hex: "#9FC3E9", Name: "6322823"},  // light royal blue, bright light blue
	{Hex: "#008F9B", Name: "6311437"},  // bright bluish green, dark turquoise
	{Hex: "#C870A0", Name: "6322821"},  // bright purple, dark pink
	{Hex: "#0A3463", Name: "6284584"},  // earth blue, dark blue
	{Hex: "#582A12", Name: "6284586"},  // reddish brown
	{Hex: "#958A73", Name: "6322841"},  // sand yellow, dark tan
	{Hex: "#E4ADC8", Name: "6284587"},  // light purple, trans pink
	{Hex: "#D09168", Name: "6343472"},  // nougat
	{Hex: "#AA7F2E", Name: "6238891"},  // warm gold, pearl gold
	{Hex: "#078BC9", Name: "6322824"},  // dark azur, dark azure
	{Hex: "#6C6E68", Name: "6284596"},  // dark stone grey, dark bluish gray
	{Hex: "#D9D9D9", Name: "6284592"},  // white glow, glow in dark white
	{Hex: "#36AEBF", Name: "6322819"},  // medium azur, medium azure
	{Hex: "#F5CD2F", Name: "6274740"},  // trans yellow
	{Hex: "#DF6695", Name: "6325421	"}, // glitter trans dark pink
	{Hex: "#A0A5A9", Name: "6284071"},  // medium stone grey, light bluish gray
	{Hex: "#AA7D55", Name: "6284589"},  // medium nougat
	{Hex: "#898788", Name: "6238890"},  // silver metallic, flat silver
	{Hex: "#DFEEA5", Name: "6284598"},  // spring yellowish green, yellowish green
	{Hex: "#923978", Name: "6322816"},  // bright reddish violet, magenta
	{Hex: "#A95500", Name: "6322840"},  // dark orange
	{Hex: "#FFF03A", Name: "6343806"},  // cool yellow, bright light yellow
	{Hex: "#5A93DB", Name: "6284602"},  // medium blue
	{Hex: "#E1D5ED", Name: "6322820"},  // lavender
	{Hex: "#C91A09", Name: "6284574"},  // bright red, trans red
	{Hex: "#0055BF", Name: "6284575"},  // bright blue, blue
	{Hex: "#184632", Name: "6396247"},  // dark green
	{Hex: "#AC78BA", Name: "6346374"},  // medium lavender
	{Hex: "#AEEFEC", Name: "6274739"},  // trans light blue, trans blue opal
	{Hex: "#E4CD9E", Name: "6284573"},  // brick yellow, tan
	{Hex: "#F6D7B3", Name: "6315196"},  // light nougat
	{Hex: "#05131D", Name: "6284070"},  // black
	{Hex: "#F08F1C", Name: "6274747"},  // trans orange
	{Hex: "#4B9F4A", Name: "6353793"},  // bright green
	{Hex: "#F2CD37", Name: "6284577"},  // bright yellow, yellow
}
