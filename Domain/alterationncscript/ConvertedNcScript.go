package alterationncscript

import (
	"fmt"
	"regexp"
)

type ConvertedNcScript struct{}

func NewConvertedNcScript() *ConvertedNcScript {
	return &ConvertedNcScript{}
}

func (c *ConvertedNcScript) Convert(source []string) ([]string, error) {
	if len(source) == 0 {
		return nil, fmt.Errorf("変換対象がありません")
	}

	isHoleSource := c.isHoleSource(source)
	isReamerSource := c.isReamerSource(source)

	var res []string
	regPercent := regexp.MustCompile(`^%$`)
	regFdNo := regexp.MustCompile(`^O\d{4}$`)
	regTool := regexp.MustCompile(`^\(T[1234]\d\)$`)
	regSpindle := regexp.MustCompile(`^\(S\d{4}\)$`)
	regG82 := regexp.MustCompile(`^\(G82\)$`)
	regG83 := regexp.MustCompile(`^\(G83\)$`)
	regG85 := regexp.MustCompile(`^\(G85\)$`)
	regX0Y0 := regexp.MustCompile(`^X0\.Y0\.$`)
	regM99 := regexp.MustCompile(`^M99$`)
	regM30 := regexp.MustCompile(`^M30$`)
	if isReamerSource {
		res = append(res, "M00")
	}
	for i := range source {
		if regPercent.MatchString(source[i]) {
			res = append(res, "")
		} else if regFdNo.MatchString(source[i]) {
			res = append(res, "("+source[i]+")")
		} else if regTool.MatchString(source[i]) {
			r := regexp.MustCompile(`\d{2}`)
			toolNums := r.FindAllStringSubmatch(source[i], 1)
			res = append(res, "T"+toolNums[0][0])
			res = append(res, "M6 Q0")
			res = append(res, "G91G0G28Z0")
			res = append(res, "G54")
			res = append(res, "G90G0X0Y0")
			res = append(res, "G0B0C0")
			res = append(res, "G0W0")
			res = append(res, "G43Z100.H"+toolNums[0][0])
			res = append(res, "M01")
		} else if regSpindle.MatchString(source[i]) {
			r := regexp.MustCompile(`S\d{4}`)
			spindle := r.FindAllStringSubmatch(source[i], 1)
			res = append(res, spindle[0][0]+"M3")
			res = append(res, "M8")
			if !isHoleSource {
				res = append(res, "G05.1Q1")
			}
		} else if regG82.MatchString(source[i]) {
			res = append(res, "G98G82R2.0Z-1.0Q2.0P500F180L0")
		} else if regG83.MatchString(source[i]) {
			res = append(res, "G98G83R2.0 Z-45.Q2.0F180L0")
		} else if regG85.MatchString(source[i]) {
			res = append(res, "G98G85R2.0 Z-35.F150L0")
		} else if regX0Y0.MatchString(source[i]) {
			res = append(res, source[i])
			// 次の行が"M99"の場合
			if isHoleSource && len(source) > i && regM99.MatchString(source[i+1]) {
				res = append(res, "M5")
				res = append(res, "M9")
				res = append(res, "G91G0G28Z0")
			}
		} else if regM99.MatchString(source[i]) {
			if isHoleSource {
				res = append(res, "(M99)")
			} else {
				res = append(res, "G05.1Q0")
				res = append(res, "M5")
				res = append(res, "M9")
				res = append(res, "G91G0G28Z0")
				res = append(res, "(M99)")
			}
		} else if regM30.MatchString(source[i]) {
			res = append(res, "G91G0G28Z0")
			res = append(res, "G91G0G28B0C0")
			res = append(res, "(M30)")
		} else {
			res = append(res, source[i])
		}
	}

	return res, nil
}

/* 穴あけのスクリプトか判定する */
func (c *ConvertedNcScript) isHoleSource(source []string) bool {
	reg := regexp.MustCompile(`^\(G8[235]\)$`)
	for i := range source {
		if reg.MatchString(source[i]) {
			return true
		}
	}

	return false
}

/* リーマのスクリプトか判定する */
func (c *ConvertedNcScript) isReamerSource(source []string) bool {
	reg := regexp.MustCompile(`^\(T15\)$`)
	for i := range source {
		if reg.MatchString(source[i]) {
			return true
		}
	}

	return false
}
