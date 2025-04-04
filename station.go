package ducky

type Region string

const (
	RegionKlangValley Region = "klang_valley"
	RegionKuching     Region = "kuching"
)

type LineType string

const (
	LineTypeLRT LineType = "lrt"
	LineTypeMR  LineType = "mr"
	LineTypeMRT LineType = "mrt"
)

type Line struct {
	Type        LineType  `json:"type"`
	ID          string    `json:"id"`
	MotisPrefix string    `json:"motis_prefix"`
	Name        string    `json:"name"`
	Region      Region    `json:"region"`
	Asset       Asset     `json:"asset"`
	Stations    []Station `json:"stations"`
}

type Asset struct {
	Owner    string `json:"owner"`
	Operator string `json:"operator"`
}

type Station struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	StationNamingRights string   `json:"station_naming_rights"`
	Lat                 float64  `json:"lat"`
	Lng                 float64  `json:"lng"`
	InterchangeStations []string `json:"interchange_stations"`
	ConnectingStations  []string `json:"connecting_stations"`
	Places              []Place  `json:"places"`
	IsTerminal          bool     `json:"is_terminal"`
}

func (s Station) DisplayName() string {
	if s.StationNamingRights == "" {
		return s.StationNamingRights
	}

	return s.Name
}

type Place struct {
	Type string  `json:"type"`
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
}

var Assets []Asset = []Asset{
	{
		Owner:    "Prasarana Malaysia Berhad",
		Operator: "Rapid Rail Sdn Bhd",
	},
	{
		Owner:    "Mass Rapid Transit Corporation Sdn Bhd",
		Operator: "Rapid Rail Sdn Bhd",
	},
}

var Lines []Line = []Line{
	{
		Type:        LineTypeLRT,
		ID:          "AG",
		MotisPrefix: "my-rail-kl",
		Name:        "Ampang",
		Region:      RegionKlangValley,
		Asset:       Assets[0],
		Stations: []Station{
			{ID: "AG18", Name: "Ampang", Lat: 3.150318, Lng: 101.760049, InterchangeStations: []string{"KG22"}, IsTerminal: true},
			{ID: "AG17", Name: "Cahaya", Lat: 3.140575, Lng: 101.756677},
			{ID: "AG16", Name: "Cempaka", Lat: 3.138324, Lng: 101.752979},
			{ID: "AG15", Name: "Pandan Indah", Lat: 3.134581, Lng: 101.746509},
			{ID: "AG14", Name: "Pandan Jaya", Lat: 3.130141, Lng: 101.739122},
			{ID: "AG13", Name: "Maluri", Lat: 3.12329, Lng: 101.727283, Places: []Place{{Type: "mall", Name: "Sunway Velocity Mall"}}},
			{ID: "AG12", Name: "Miharja", Lat: 3.120973, Lng: 101.717922, Places: []Place{{Type: "mall", Name: "Viva Mall"}}},
			{ID: "AG11", Name: "Chan Sow Lin", Lat: 3.128105, Lng: 101.715637, InterchangeStations: []string{"PY24", "SP11"}, Places: []Place{{Type: "mall", Name: "The Metro Mall"}}},
			{ID: "AG10", Name: "Pudu", Lat: 3.134879, Lng: 101.711957, InterchangeStations: []string{"SP10"}},
			{ID: "AG9", Name: "Hang Tuah", StationNamingRights: "BBCC Hang Tuah", Lat: 3.140012, Lng: 101.705984, InterchangeStations: []string{"SP9", "MR4"}, Places: []Place{{Type: "mall", Name: "Berjaya Times Square"}, {Type: "mall", Name: "The Mitsui Shopping Park LaLaport Bukit Bintang City Centre"}, {Type: "mosque", Name: "Masjid Al Bukhari", Lat: 3.1391978, Lng: 101.7046318}}},
			{ID: "AG8", Name: "Plaza Rakyat", Lat: 3.144049, Lng: 101.702105, InterchangeStations: []string{"SP8", "KG17"}},
			{ID: "AG7", Name: "Masjid Jamek", Lat: 3.14927, Lng: 101.696377, InterchangeStations: []string{"KJ13", "SP7"}},
			{ID: "AG6", Name: "Bandaraya", Lat: 3.155567, Lng: 101.694485, InterchangeStations: []string{"SP6"}, Places: []Place{{Type: "mall", Name: "SOGO Kuala Lumpur"}}},
			{ID: "AG5", Name: "Sultan Ismail", Lat: 3.161245, Lng: 101.694109, InterchangeStations: []string{"SP5"}, ConnectingStations: []string{"MR9"}},
			{ID: "AG4", Name: "PWTC", Lat: 3.166333, Lng: 101.693586, InterchangeStations: []string{"SP4"}, Places: []Place{{Type: "mall", Name: "Sunway Putra Mall"}}},
			{ID: "AG3", Name: "Titiwangsa", Lat: 3.173497, Lng: 101.695367, InterchangeStations: []string{"PY17", "SP3", "MR11"}},
			{ID: "AG2", Name: "Sentul", Lat: 3.178484, Lng: 101.695542, InterchangeStations: []string{"SP2"}},
			{ID: "AG1", Name: "Sentul Timur", Lat: 3.185897, Lng: 101.695217, InterchangeStations: []string{"SP1"}, IsTerminal: true},
		},
	},
	{
		Type:        "LRT",
		ID:          "SP",
		MotisPrefix: "my-rail-kl",
		Name:        "Sri Petaling",
		Region:      RegionKlangValley,
		Asset:       Assets[0],
		Stations: []Station{
			{ID: "SP31", Name: "Putra Heights", Lat: 2.996016, Lng: 101.575521, InterchangeStations: []string{"KJ37"}, Places: []Place{{Type: "mosque", Name: "Masjid Putra Height", Lat: 2.9975838904405965, Lng: 101.57626199789122}}, IsTerminal: true},
			{ID: "SP29", Name: "Puchong Prima", Lat: 2.999808, Lng: 101.596692},
			{ID: "SP28", Name: "Puchong Perdana", Lat: 3.007913, Lng: 101.605021},
			{ID: "SP27", Name: "Bandar Puteri", Lat: 3.017111, Lng: 101.612855},
			{ID: "SP26", Name: "Taman Perindustrian Puchong", Lat: 3.022814, Lng: 101.613514},
			{ID: "SP25", Name: "Pusat Bandar Puchong", Lat: 3.033194, Lng: 101.616057, Places: []Place{{Type: "mall", Name: "SetiaWalk"}}},
			{ID: "SP24", Name: "IOI Puchong Jaya", Lat: 3.048101, Lng: 101.62095, Places: []Place{{Type: "mall", Name: "IOI Mall Puchong"}}},
			{ID: "SP22", Name: "Kinrara", Lat: 3.050506, Lng: 101.644294},
			{ID: "SP21", Name: "Alam Sutera", Lat: 3.0547, Lng: 101.656468},
			{ID: "SP20", Name: "Muhibbah", Lat: 3.062229, Lng: 101.662552},
			{ID: "SP19", Name: "Awan Besar", Lat: 3.062131, Lng: 101.670555},
			{ID: "SP18", Name: "Sri Petaling", Lat: 3.061445, Lng: 101.687074},
			{ID: "SP17", Name: "Bukit Jalil", Lat: 3.058196, Lng: 101.692125, Places: []Place{{Type: "mall", Name: "Endah Parade"}}},
			{ID: "SP16", Name: "Sungai Besi", Lat: 3.063842, Lng: 101.708062, InterchangeStations: []string{"PY29"}, Places: []Place{{Type: "mosque", Name: "Masjid Jamek Sungai Besi (Ibnu Khaldun)", Lat: 3.064344710399052, Lng: 101.70925452938673}}},
			{ID: "SP15", Name: "Bandar Tasik SeLatan", Lat: 3.076058, Lng: 101.711107},
			{ID: "SP14", Name: "Bandar Tun Razak", Lat: 3.089576, Lng: 101.712466},
			{ID: "SP13", Name: "Salak SeLatan", Lat: 3.102201, Lng: 101.706179},
			{ID: "SP12", Name: "Cheras", Lat: 3.112609, Lng: 101.714178},
			{ID: "SP11", Name: "Chan Sow Lin", Lat: 3.128105, Lng: 101.715637, InterchangeStations: []string{"PY24", "AG11"}, Places: []Place{{Type: "mall", Name: "The Metro Mall"}}},
			{ID: "SP10", Name: "Pudu", Lat: 3.134879, Lng: 101.711957, InterchangeStations: []string{"AG10"}},
			{ID: "SP9", Name: "Hang Tuah", StationNamingRights: "BBCC Hang Tuah", Lat: 3.140012, Lng: 101.705984, InterchangeStations: []string{"MR4", "AG9"}, Places: []Place{{Type: "mall", Name: "Berjaya Times Square"}, {Type: "mall", Name: "The Mitsui Shopping Park LaLaport Bukit Bintang City Centre"}, {Type: "mosque", Name: "Masjid Al Bukhari", Lat: 3.1391978, Lng: 101.7046318}}},
			{ID: "SP8", Name: "Plaza Rakyat", Lat: 3.144049, Lng: 101.702105, InterchangeStations: []string{"AG8", "KG17"}},
			{ID: "SP7", Name: "Masjid Jamek", Lat: 3.14927, Lng: 101.696377, InterchangeStations: []string{"AG7", "KJ13"}},
			{ID: "SP6", Name: "Bandaraya", Lat: 3.155567, Lng: 101.694485, InterchangeStations: []string{"AG6"}, Places: []Place{{Type: "mall", Name: "SOGO Kuala Lumpur"}}},
			{ID: "SP5", Name: "Sultan Ismail", Lat: 3.161245, Lng: 101.694109, InterchangeStations: []string{"AG5"}, ConnectingStations: []string{"MR9"}},
			{ID: "SP4", Name: "PWTC", Lat: 3.166333, Lng: 101.693586, InterchangeStations: []string{"AG4"}, Places: []Place{{Type: "mall", Name: "Sunway Putra Mall"}}},
			{ID: "SP3", Name: "Titiwangsa", Lat: 3.173497, Lng: 101.695367, InterchangeStations: []string{"PY17", "AG3", "MR11"}},
			{ID: "SP2", Name: "Sentul", Lat: 3.178484, Lng: 101.695542, InterchangeStations: []string{"AG2"}},
			{ID: "SP1", Name: "Sentul Timur", Lat: 3.185897, Lng: 101.695217, InterchangeStations: []string{"AG1"}, IsTerminal: true},
		},
	},
	{
		Type:        LineTypeLRT,
		ID:          "KJ",
		MotisPrefix: "my-rail-kl",
		Name:        "Kelana Jaya",
		Region:      RegionKlangValley,
		Asset:       Assets[0],
		Stations: []Station{
			{ID: "KJ37", Name: "Putra Heights", Lat: 2.996227, Lng: 101.575462, InterchangeStations: []string{"SP31"}, Places: []Place{{Type: "mosque", Name: "Masjid Putra Height", Lat: 2.9975838904405965, Lng: 101.57626199789122}}, IsTerminal: true},
			{ID: "KJ36", Name: "Subang Alam", Lat: 3.009421, Lng: 101.572281},
			{ID: "KJ35", Name: "Alam Megah", Lat: 3.023151, Lng: 101.572029},
			{ID: "KJ34", Name: "USJ 21", Lat: 3.029881, Lng: 101.581711, Places: []Place{{Type: "mall", Name: "The Place Mall"}, {Type: "mosque", Name: "Masjid Al-Madaniah", Lat: 3.0313921849430185, Lng: 101.58398391263445}}},
			{ID: "KJ33", Name: "Wawasan", Lat: 3.035062, Lng: 101.588348, Places: []Place{{Type: "mall", Name: "The 19 USJ City Mall (Palazzo 19 Mall)"}}},
			{ID: "KJ32", Name: "Taipan", Lat: 3.04815, Lng: 101.590233, Places: []Place{{Type: "mosque", Name: "Masjid Al-Falah USJ 9", Lat: 3.0440662790664836, Lng: 101.58719996488976}}},
			{ID: "KJ31", Name: "USJ 7", Lat: 3.054956, Lng: 101.592194, Places: []Place{{Type: "mall", Name: "DA MEN Mall"}}},
			{ID: "KJ30", Name: "SS 18", Lat: 3.067182, Lng: 101.585945},
			{ID: "KJ29", Name: "SS 15", Lat: 3.075972, Lng: 101.585983, Places: []Place{{Type: "mall", Name: "SS15 Courtyard"}, {Type: "mosque", Name: "Masjid Darul Ehsan Subang Jaya", Lat: 3.0804905148779658, Lng: 101.58554713022856}}},
			{ID: "KJ28", Name: "Subang Jaya", Lat: 3.08466, Lng: 101.588127, Places: []Place{{Type: "mall", Name: "NU Empire Shopping Gallery"}, {Type: "mall", Name: "Subang Parade Shopping Centre"}, {Type: "mall", Name: "AEON BiG Subang Jaya"}}},
			{ID: "KJ27", Name: "Glenmarie", StationNamingRights: "CGC Glenmarie", Lat: 3.094732, Lng: 101.590622},
			{ID: "KJ26", Name: "Ara Damansara", Lat: 3.108643, Lng: 101.586372, Places: []Place{{Type: "mall", Name: "Evolve Concept Mall"}}},
			{ID: "KJ25", Name: "Lembah Subang", Lat: 3.112094, Lng: 101.591034},
			{ID: "KJ24", Name: "Kelana Jaya", Lat: 3.112497, Lng: 101.6043},
			{ID: "KJ23", Name: "Taman Bahagia", Lat: 3.11079, Lng: 101.612856},
			{ID: "KJ22", Name: "Taman Paramount", Lat: 3.104716, Lng: 101.623192},
			{ID: "KJ21", Name: "Asia Jaya", Lat: 3.104343, Lng: 101.637695},
			{ID: "KJ20", Name: "Taman Jaya", Lat: 3.104086, Lng: 101.645248, Places: []Place{{Type: "mall", Name: "Amcorp Mall"}}},
			{ID: "KJ19", Name: "Universiti", Lat: 3.114616, Lng: 101.661639, Places: []Place{{Type: "mall", Name: "KL Gateway Mall"}, {Type: "mosque", Name: "Masjid Ar-Rahman", Lat: 3.1178179, Lng: 101.6628105}}},
			{ID: "KJ18", Name: "Kerinchi", Lat: 3.115506, Lng: 101.668572, Places: []Place{
				{Type: "mosque", Name: "Masjid Ar-Rahah", Lat: 3.1137199, Lng: 101.6687277},
				{Type: "mosque", Name: "Masjid Menara TM", Lat: 3.1161164, Lng: 101.6656948},
			}},
			{ID: "KJ17", Name: "Abdullah Hukum", Lat: 3.118735, Lng: 101.672897, Places: []Place{{Type: "mall", Name: "Mid Valley Megamall"}, {Type: "mall", Name: "The Gardens Mall"}, {Type: "mosque", Name: "Masjid TNB", Lat: 3.1177874, Lng: 101.6709643}}},
			{ID: "KJ16", Name: "Bangsar", StationNamingRights: "Bank Rakyat Bangsar", Lat: 3.127588, Lng: 101.679062},
			{ID: "KJ15", Name: "KL Sentral", StationNamingRights: "KL Sentral redONE", Lat: 3.13442, Lng: 101.68625, ConnectingStations: []string{"MR1"}, Places: []Place{{Type: "mall", Name: "NU Sentral"}}},
			{ID: "KJ14", Name: "Pasar Seni", Lat: 3.142439, Lng: 101.69531, InterchangeStations: []string{"PY14"}, Places: []Place{{Type: "mosque", Name: "Masjid Negara", Lat: 3.1419713907686377, Lng: 101.69174639937577}}},
			{ID: "KJ13", Name: "Masjid Jamek", Lat: 3.149714, Lng: 101.696815, InterchangeStations: []string{"AG7", "SP7"}, Places: []Place{{Type: "mosque", Name: "Masjid Jamek Sultan Abdul Samad", Lat: 3.1489147, Lng: 101.695355}}},
			{ID: "KJ12", Name: "Dang Wangi", Lat: 3.156942, Lng: 101.701975, ConnectingStations: []string{"MR8"}},
			{ID: "KJ11", Name: "Kampung Baru", Lat: 3.161386, Lng: 101.706608, Places: []Place{{Type: "mosque", Name: "Masjid Jamek Kg Baru", Lat: 3.1642272, Lng: 101.7012624}}},
			{ID: "KJ10", Name: "KLCC", Lat: 3.158935, Lng: 101.713287, Places: []Place{{Type: "landmark", Name: "Petronas Twin Towers"}, {Type: "mall", Name: "Suria KLCC"}, {Type: "mall", Name: "Avenue K"}, {Type: "mosque", Name: "Masjid As-Syakirin", Lat: 3.1572322, Lng: 101.713708}}},
			{ID: "KJ9", Name: "Ampang Park", Lat: 3.159894, Lng: 101.719017, ConnectingStations: []string{"PY20"}, Places: []Place{{Type: "mall", Name: "The LINC KL"}, {Type: "mall", Name: "Intermark Mall"}}},
			{ID: "KJ8", Name: "Damai", Lat: 3.164406, Lng: 101.724489},
			{ID: "KJ7", Name: "Dato' Keramat", Lat: 3.16509, Lng: 101.73184, Places: []Place{{Type: "mosque", Name: "Masjid Al-Akram Datuk Keramat", Lat: 3.1664534702473905, Lng: 101.72950400592192}}},
			{ID: "KJ6", Name: "Jelatek", Lat: 3.167204, Lng: 101.735344, Places: []Place{{Type: "mall", Name: "Datum Jelatek Shopping Centre"}}},
			{ID: "KJ5", Name: "Setiawangsa", Lat: 3.17576, Lng: 101.73584, Places: []Place{{Type: "mosque", Name: "Masjid Muadz bin Jabal", Lat: 3.1779773, Lng: 101.7361406}}},
			{ID: "KJ4", Name: "Sri Rampai", Lat: 3.199176, Lng: 101.73747, Places: []Place{{Type: "mall", Name: "Wangsa Walk Mall"}}},
			{ID: "KJ3", Name: "Wangsa Maju", Lat: 3.205751, Lng: 101.731796, Places: []Place{{Type: "mosque", Name: "Masjid Usamah Bin Zaid", Lat: 3.2029713366589996, Lng: 101.73692327782841}}},
			{ID: "KJ2", Name: "Taman Melati", Lat: 3.219558, Lng: 101.72197, Places: []Place{{Type: "mall", Name: "M3 Shopping Mall"}, {Type: "mosque", Name: "Masjid Salahudin Al-Ayyubi", Lat: 3.2225904, Lng: 101.7177032}}},
			{ID: "KJ1", Name: "Gombak", Lat: 3.231793, Lng: 101.724427, IsTerminal: true},
		},
	},
	{
		Type:        LineTypeMR,
		ID:          "MR",
		MotisPrefix: "my-rail-kl",
		Name:        "Monorail KL",
		Region:      RegionKlangValley,
		Asset:       Assets[0],
		Stations: []Station{
			{ID: "MR1", Name: "KL Sentral", Lat: 3.132852, Lng: 101.687817, ConnectingStations: []string{"KJ15"}, Places: []Place{{Type: "mall", Name: "NU Sentral"}}},
			{ID: "MR2", Name: "Tun Sambanthan", Lat: 3.13132, Lng: 101.69085},
			{ID: "MR3", Name: "Maharajalela", Lat: 3.138743, Lng: 101.699268, Places: []Place{{Type: "mosque", Name: "Masjid Al-Sultan Abdullah", Lat: 3.1397038347130954, Lng: 101.69941427658394}}},
			{ID: "MR4", Name: "Hang Tuah", Lat: 3.140511, Lng: 101.706029, InterchangeStations: []string{"AG9", "SP9"}, Places: []Place{{Type: "mall", Name: "Berjaya Times Square"}, {Type: "mall", Name: "The Mitsui Shopping Park LaLaport Bukit Bintang City Centre"}, {Type: "mosque", Name: "Masjid Al Bukhari", Lat: 3.1391978, Lng: 101.7046318}}},
			{ID: "MR5", Name: "Imbi", Lat: 3.14283, Lng: 101.70945, Places: []Place{{Type: "mall", Name: "Berjaya Times Square"}}},
			{ID: "MR6", Name: "Bukit Bintang", Lat: 3.146022, Lng: 101.7115, ConnectingStations: []string{"KG18A"}},
			{ID: "MR7", Name: "Raja Chulan", Lat: 3.150878, Lng: 101.710432},
			{ID: "MR8", Name: "Bukit Nanas", Lat: 3.156214, Lng: 101.704809, ConnectingStations: []string{"KJ12"}},
			{ID: "MR9", Name: "Medan Tuanku", Lat: 3.15935, Lng: 101.69888, ConnectingStations: []string{"AG5", "SP5"}, Places: []Place{{Type: "mall", Name: "Quill City Mall Kuala Lumpur"}}},
			{ID: "MR10", Name: "Chow Kit", Lat: 3.167358, Lng: 101.698379},
			{ID: "MR11", Name: "Titiwangsa", Lat: 3.173192, Lng: 101.696022, InterchangeStations: []string{"AG3", "SP3", "PY17"}},
		},
	},
	{
		Type:        LineTypeMRT,
		ID:          "KG",
		MotisPrefix: "my-rail-kl",
		Name:        "Kajang",
		Region:      RegionKlangValley,
		Asset:       Assets[1],
		Stations: []Station{
			{ID: "KG04", Name: "Kwasa Damansara", Lat: 3.176146, Lng: 101.572052, InterchangeStations: []string{"PY01"}, IsTerminal: true},
			{ID: "KG05", Name: "Kwasa Sentral", Lat: 3.170112, Lng: 101.564651},
			{ID: "KG06", Name: "Kota Damansara", Lat: 3.150134, Lng: 101.57869},
			{ID: "KG07", Name: "Surian", Lat: 3.14948, Lng: 101.593925, Places: []Place{{Type: "mall", Name: "IOI Mall Damansara"}, {Type: "mall", Name: "Sunway Giza Mall"}}},
			{ID: "KG08", Name: "Mutiara Damansara", Lat: 3.155301, Lng: 101.609077, Places: []Place{{Type: "mall", Name: "The Curve"}, {Type: "mall", Name: "IPC Shopping Centre"}, {Type: "store", Name: "IKEA Damansara"}}},
			{ID: "KG09", Name: "Bandar Utama", Lat: 3.14671, Lng: 101.618599, Places: []Place{{Type: "mall", Name: "1 Utama Shopping Centre"}}},
			{ID: "KG10", Name: "Taman Tun Dr Ismail", StationNamingRights: "TTDI Deloitte", Lat: 3.13613, Lng: 101.630539, Places: []Place{{Type: "mall", Name: "Glo Damansara Shopping Mall"}}},
			{ID: "KG12", Name: "Phileo Damansara", Lat: 3.129864, Lng: 101.642471},
			{ID: "KG13", Name: "Pusat Bandar Damansara", StationNamingRights: "Pavilion Pusat Bandar Damansara", Lat: 3.143444, Lng: 101.662857, Places: []Place{{Type: "mall", Name: "Pavilion Damansara Heights"}}},
			{ID: "KG14", Name: "Semantan", Lat: 3.150977, Lng: 101.665497},
			{ID: "KG15", Name: "Muzium Negara", Lat: 3.137317, Lng: 101.687336, ConnectingStations: []string{"KJ15"}, Places: []Place{{Type: "mall", Name: "NU Sentral"}}},
			{ID: "KG16", Name: "Pasar Seni", Lat: 3.142293265, Lng: 101.6955642, InterchangeStations: []string{"KJ14"}, Places: []Place{{Type: "mosque", Name: "Masjid Negara", Lat: 3.1419713907686377, Lng: 101.69174639937577}}},
			{ID: "KG17", Name: "Merdeka", Lat: 3.141969, Lng: 101.70205, InterchangeStations: []string{"AG8", "SP8"}, Places: []Place{{Type: "mosque", Name: "Masjid Al-Sultan Abdullah", Lat: 3.1395985, Lng: 101.6994226}}},
			{
				ID: "KG18A", Name: "Bukit Bintang", StationNamingRights: "Pavilion Kuala Lumpur Bukit Bintang", Lat: 3.146503, Lng: 101.710947, ConnectingStations: []string{"MR6"}, Places: []Place{
					{Type: "mall", Name: "Pavilion Kuala Lumpur"},
					{Type: "mall", Name: "Fahrenheit88"},
					{Type: "mall", Name: "Lot 10 Shopping Centre"},
					{Type: "mall", Name: "Low Yat Plaza"},
					{Type: "mall", Name: "Sungei Wang Plaza"},
					{Type: "mall", Name: "The Starhill"},
				},
			},
			{
				ID: "KG20", Name: "Tun Razak Exchange", StationNamingRights: "Tun Razak Exchange Samsung Galaxy", Lat: 3.142403, Lng: 101.720156, InterchangeStations: []string{"PY23"}, Places: []Place{
					{Type: "mall", Name: "The Exchange TRX"},
					{Type: "store", Name: "Apple The Exchange TRX"},
				},
			},
			{
				ID: "KG21", Name: "Cochrane", Lat: 3.132829, Lng: 101.722962, Places: []Place{
					{Type: "mall", Name: "MyTOWN Shopping Centre"},
					{Type: "store", Name: "IKEA Cheras"},
					{Type: "mall", Name: "Sunway Velocity Mall"},
					{Type: "mosque", Name: "Masjid Jamek Alam Shah", Lat: 3.1350761, Lng: 101.7180825},
				},
			},
			{ID: "KG22", Name: "Maluri", StationNamingRights: "AEON Maluri", Lat: 3.123623, Lng: 101.727809, InterchangeStations: []string{"AG13"}, Places: []Place{{Type: "mall", Name: "Sunway Velocity Mall"}}},
			{ID: "KG23", Name: "Taman Pertama", Lat: 3.112547, Lng: 101.729371},
			{ID: "KG24", Name: "Taman Midah", Lat: 3.104505, Lng: 101.732186},
			{
				ID: "KG25", Name: "Taman Mutiara", Lat: 3.090989, Lng: 101.740453, Places: []Place{
					{Type: "mall", Name: "EkoCheras Mall"},
					{Type: "mall", Name: "Cheras LeisureMall"},
				},
			},
			{ID: "KG26", Name: "Taman Connaught", Lat: 3.079172, Lng: 101.74522, Places: []Place{{Type: "mall", Name: "Cheras Sentral Mall"}}},
			{ID: "KG27", Name: "Taman Suntex", Lat: 3.071578, Lng: 101.763552},
			{ID: "KG28", Name: "Sri Raya", Lat: 3.062273, Lng: 101.772899},
			{ID: "KG29", Name: "Bandar Tun Hussien Onn", Lat: 3.048223, Lng: 101.775109},
			{ID: "KG30", Name: "Batu 11 Cheras", Lat: 3.041339, Lng: 101.773383},
			{ID: "KG31", Name: "Bukit Dukung", Lat: 3.026413, Lng: 101.771072},
			{ID: "KG33", Name: "Sungai Jernih", Lat: 3.000948, Lng: 101.783857},
			{ID: "KG34", Name: "Stadium Kajang", Lat: 2.994514, Lng: 101.786338, Places: []Place{{Type: "mall", Name: "Plaza Metro Kajang"}}},
			{ID: "KG35", Name: "Kajang", Lat: 2.982778, Lng: 101.790278, IsTerminal: true},
		},
	},
	{
		Type:        LineTypeMRT,
		ID:          "PY",
		MotisPrefix: "my-rail-kl",
		Name:        "Putrajaya",
		Region:      RegionKlangValley,
		Asset:       Assets[1],
		Stations: []Station{
			{ID: "PY01", Name: "Kwasa Damansara", Lat: 3.1763324, Lng: 101.5721456, InterchangeStations: []string{"KG04"}, IsTerminal: true},
			{ID: "PY03", Name: "Kampung Selamat", Lat: 3.197266, Lng: 101.578499, Places: []Place{{Type: "gym", Name: "Anytime Fitness SqWhere"}}},
			{ID: "PY04", Name: "Sungai Buloh", Lat: 3.206429, Lng: 101.581779},
			{ID: "PY05", Name: "Damansara Damai", Lat: 3.199892, Lng: 101.592623},
			{ID: "PY06", Name: "Sri Damansara Barat", Lat: 3.198197, Lng: 101.608302, Places: []Place{{Type: "gym", Name: "Anytime Fitness Sri Damansara"}}},
			{ID: "PY07", Name: "Sri Damansara Sentral", Lat: 3.198815, Lng: 101.621396},
			{ID: "PY08", Name: "Sri Damansara Timur", Lat: 3.207832, Lng: 101.628716, Places: []Place{{Type: "mall", Name: "Kompleks Desa Kepong"}}},
			{ID: "PY09", Name: "Metro Prima", Lat: 3.214438, Lng: 101.639402},
			{ID: "PY10", Name: "Kepong Baru", Lat: 3.211663, Lng: 101.648193},
			{ID: "PY11", Name: "Jinjang", Lat: 3.209544, Lng: 101.655829},
			{ID: "PY12", Name: "Sri Delima", Lat: 3.207108, Lng: 101.665749, Places: []Place{{Type: "mall", Name: "Brem Mall Shopping Complex"}}},
			{ID: "PY13", Name: "Kampung Batu", Lat: 3.205521, Lng: 101.675473},
			{ID: "PY14", Name: "Kentomen", Lat: 3.19563, Lng: 101.6797},
			{ID: "PY15", Name: "Jalan Ipoh", Lat: 3.189319, Lng: 101.681145},
			{ID: "PY16", Name: "Sentul Barat", Lat: 3.179369, Lng: 101.684742},
			{ID: "PY17", Name: "Titiwangsa", Lat: 3.17408, Lng: 101.69581, InterchangeStations: []string{"AG3", "SP3", "MR11"}},
			{ID: "PY18", Name: "Hospital Kuala Lumpur", Lat: 3.17405, Lng: 101.70239, Places: []Place{{Type: "hospital", Name: "Hospital Kuala Lumpur"}, {Type: "hospital", Name: "KPJ Tawakkal KL Specialist Hospital"}}},
			{ID: "PY19", Name: "Raja Uda", StationNamingRights: "Raja Uda UTM", Lat: 3.16794, Lng: 101.71017},
			{ID: "PY20", Name: "Ampang Park", Lat: 3.16225, Lng: 101.71781, ConnectingStations: []string{"KJ9"}},
			{ID: "PY21", Name: "Persiaran KLCC", Lat: 3.15712, Lng: 101.71834},
			{ID: "PY22", Name: "Conlay", StationNamingRights: "Conlay Kompleks Kraf", Lat: 3.15145, Lng: 101.71801},
			{ID: "PY23", Name: "Tun Razak Exchange", Lat: 3.14289, Lng: 101.72034, InterchangeStations: []string{"KG20"}, Places: []Place{{Type: "mall", Name: "The Exchange TRX"}, {Type: "store", Name: "Apple The Exchange TRX"}}},
			{ID: "PY24", Name: "Chan Sow Lin", Lat: 3.12839, Lng: 101.71663, InterchangeStations: []string{"AG11", "SP11"}, Places: []Place{{Type: "mall", Name: "The Metro Mall"}}},
			{ID: "PY27", Name: "Kuchai", Lat: 3.089546, Lng: 101.694124},
			{ID: "PY28", Name: "Taman Naga Emas", Lat: 3.077688, Lng: 101.699867},
			{ID: "PY29", Name: "Sungai Besi", Lat: 3.063737, Lng: 101.7084, InterchangeStations: []string{"SP16"}, Places: []Place{{Type: "mosque", Name: "Masjid Jamek Sungai Besi (Ibnu Khaldun)", Lat: 3.064344710399052, Lng: 101.70925452938673}}},
			{ID: "PY31", Name: "Serdang Raya Utara", Lat: 3.041674, Lng: 101.704928, Places: []Place{{Type: "mosque", Name: "Masjid Al-Islah Serdang Raya", Lat: 3.043026344004756, Lng: 101.70343677890706}}},
			{ID: "PY32", Name: "Serdang Raya Selatan", Lat: 3.028463, Lng: 101.707514},
			{ID: "PY33", Name: "Serdang Jaya", Lat: 3.0216, Lng: 101.709},
			{ID: "PY34", Name: "UPM", Lat: 3.008489, Lng: 101.705396},
			{ID: "PY36", Name: "Taman Equine", Lat: 2.98942, Lng: 101.67244},
			{ID: "PY37", Name: "Putra Permai", Lat: 2.98339, Lng: 101.66099},
			{ID: "PY38", Name: "16 Sierra", Lat: 2.964974, Lng: 101.654812},
			{ID: "PY39", Name: "Cyberjaya Utara", StationNamingRights: "Cyberjaya Utara Finexus", Lat: 2.95, Lng: 101.6573},
			{ID: "PY40", Name: "Cyberjaya City Centre", StationNamingRights: "Cyberjaya City Centre Limkokwing", Lat: 2.9384, Lng: 101.6659},
			{ID: "PY41", Name: "Putrajaya Sentral", Lat: 2.9313, Lng: 101.6715, Places: []Place{{Type: "terminal", Name: "Terminal Putrajaya Sentral"}}, IsTerminal: true},
		},
	},
}
