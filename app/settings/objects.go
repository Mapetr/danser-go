package settings

var Objects = initObjects()

func initObjects() *objects {
	return &objects{
		ForceSliderBallTexture: true,
		DrawApproachCircles:    true,
		DrawComboNumbers:       true,
		DrawReverseArrows:      true,
		DrawSliderFollowCircle: true,
		DrawFollowPoints:       true,
		LoadSpinners:           false,
		ScaleToTheBeat:         false,
		SliderLOD:              50,
		SliderPathLOD:          50,
		SliderSnakeIn:          true,
		SliderSnakeInMult:      0.0,
		SliderSnakeOut:         true,
		SliderMerge:            true,
		SliderDistortions:      true,
		DrawScorePoints:        true,
		StackEnabled:           true,
		Colors: &objectcolors{
			MandalaTexturesTrigger: 5,
			MandalaTexturesAlpha:   0.3,
			Color: &color{
				EnableRainbow: true,
				RainbowSpeed:  8,
				BaseColor: &hsv{
					0,
					1.0,
					1.0},
				EnableCustomHueOffset: false,
				HueOffset:             0,
				FlashToTheBeat:        false,
				FlashAmplitude:        100,
				currentHue:            0,
			},
			UseComboColors:                false,
			ComboColors:                   []*hsv{{Hue: 0, Saturation: 1, Value: 1}},
			WhiteScorePoints:              false,
			ScorePointColorOffset:         0,
			EnableCustomSliderBorderColor: false,
			CustomSliderBorderColor: &color{
				EnableRainbow: false,
				RainbowSpeed:  8,
				BaseColor: &hsv{
					0,
					0.0,
					1.0},
				EnableCustomHueOffset: false,
				HueOffset:             0,
				FlashToTheBeat:        true,
				FlashAmplitude:        100,
				currentHue:            0,
			},
			EnableCustomSliderBorderGradientOffset: true,
			SliderBorderGradientOffset:             18,
		},
	}
}

type objects struct {
	ForceSliderBallTexture bool //true, if this is disabled, mandala texture will be used for slider ball
	DrawApproachCircles    bool //true
	DrawComboNumbers       bool
	DrawReverseArrows      bool
	DrawSliderFollowCircle bool
	DrawFollowPoints       bool
	LoadSpinners           bool
	ScaleToTheBeat         bool  //true, objects size is changing with music peak amplitude
	SliderLOD              int64 //30, number of triangles in a circle
	SliderPathLOD          int64 //50, int(pixelLength*(PathLOD/100)) => number of slider path points
	SliderSnakeIn          bool
	SliderSnakeInMult      float64
	SliderSnakeOut         bool
	SliderMerge            bool
	SliderDistortions      bool //true, osu!stable slider distortions on aspire maps
	DrawScorePoints        bool //true
	StackEnabled           bool //true, stack leniency
	Colors                 *objectcolors
}

type objectcolors struct {
	MandalaTexturesTrigger                 int     //5, minimum value of cursors needed to use more translucent texture
	MandalaTexturesAlpha                   float64 //0.3
	Color                                  *color
	UseComboColors                         bool
	ComboColors                            []*hsv
	WhiteScorePoints                       bool    //true
	ScorePointColorOffset                  float64 //0.0, hue offset of the followpoint
	EnableCustomSliderBorderColor          bool
	CustomSliderBorderColor                *color
	EnableCustomSliderBorderGradientOffset bool
	SliderBorderGradientOffset             float64 //18, hue offset of slider outer border
}
