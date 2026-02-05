package converter

import (
	"github.com/kingtingthegreat/img2lego/util"
	"github.com/kyroy/kdtree"
	"github.com/kyroy/kdtree/points"
)

type Lookup struct {
	kdt *kdtree.KDTree
}

func InitLookup(colors []Color) *Lookup {
	kdt := kdtree.New([]kdtree.Point{})

	for _, c := range colors {
		rgb := util.HexToRGB(c.Hex)
		kdt.Insert(rgb.Data)
	}

	return &Lookup{
		kdt: kdt,
	}
}

func (l *Lookup) FindMostSimilar(rgb util.RGB) util.RGB {
	closest := l.kdt.KNN(rgb.Data, 1)[0].(*points.Point3D)
	return util.RGB{Data: closest}
}
