package utils

import (
	"database/sql/driver"
	"fmt"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GeoPoint struct {
	X, Y float64
}

// GormValue 将查询条件转化成SQL语句的功能
func (g GeoPoint) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_GeomFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", g.X, g.Y)},
	}
}

// Scan 解析pg原始数据
func (g *GeoPoint) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	pt, err := ewkbhex.Decode(value.(string))
	if err == nil {
		if p, ok := pt.(*geom.Point); ok {
			g.X = p.X()
			g.Y = p.Y()
		} else {
			// return errors.New(fmt.Sprint("Failed to unmarshal geometry:", pt))
			fmt.Sprint("Failed to unmarshal geometry:", pt)
			return nil
		}
	}
	return err
}

// Value gorm用于将原始数据转化为WKT
func (g GeoPoint) Value() (driver.Value, error) {
	return fmt.Sprintf("POINT(%f %f)", g.X, g.Y), nil
}

type MultiPoint geom.MultiPoint
type LineString geom.LineString

func (g GeoPoint) GeoJSONType() string {
	return "Point"
}

func (mp MultiPoint) GeoJSONType() string {
	return "MultiPoint"
}

func (l LineString) GeoJSONType() string {
	return "LineString"
}

type Geometry interface {
	GeoJSONType() string
}
