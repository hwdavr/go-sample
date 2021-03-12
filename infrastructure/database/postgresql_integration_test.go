// +build integration

package database_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/pkg/errors"

	. "github.com/hwdavr/perfmon/domain"
	"github.com/hwdavr/perfmon/infrastructure/database"
	. "github.com/hwdavr/perfmon/internal/context"
)

var (
	dbPsql = database.NewPsql("test")
)

func TestPsqlCreate(t *testing.T) {
	t.Run("", func(t *testing.T) {
		product := "p1"
		version := "v1"
		attributes := []string{"a11", "a12"}
		name := "n1"
		measurement := "m1"
		timestamp := time.Date(2018, 11, 8, 16, 23, 50, 0, time.UTC)
		value := 1
		createdOn := time.Now()
		createdBy := "192.168.1.2"
		job := &Job{product, version, attributes, name, measurement, timestamp, value}
		ctx := context.Background()
		ctx = WithCreatedOn(ctx, createdOn)
		ctx = WithCreatedBy(ctx, createdBy)
		err := dbPsql.Create(ctx, job)
		if err != nil {
			t.Errorf("error occurred: %+v", err)
		}
		jobs, err := dbPsql.Jobs(product, version, attributes, name, measurement, timestamp, timestamp)
		if err != nil {
			t.Errorf("error occurred: %+v", err)
		}
		got := len(jobs)
		want := 1
		if got != want {
			t.Errorf("len(dbPsql.Jobs(%v, %v, %v, %v, %v, %v, %v)) = %v; want %v", product, version, attributes, name, measurement, timestamp, timestamp, got, want)
		}
		err = dbPsql.Delete(job)
		if err != nil {
			t.Errorf("error occurred: %+v", err)
		}
	})
	t.Run("", func(t *testing.T) {
		product := "p6"
		version := "v6"
		attributes := []string{"a61", "a62"}
		name := "n6"
		measurement := "m6"
		timestamp := time.Date(2017, 10, 29, 16, 18, 07, 0, time.UTC)
		value := 48
		createdOn := time.Now()
		createdBy := "192.168.1.2"
		job := &Job{product, version, attributes, name, measurement, timestamp, value}
		ctx := context.Background()
		ctx = WithCreatedOn(ctx, createdOn)
		ctx = WithCreatedBy(ctx, createdBy)
		err := dbPsql.Create(ctx, job)
		if err != nil {
			t.Errorf("error occurred: %+v", err)
		}
		jobs, err := dbPsql.Jobs(product, version, attributes, name, measurement, timestamp, timestamp)
		if err != nil {
			t.Errorf("error occurred: %+v", err)
		}
		got := len(jobs)
		want := 2
		if got != want {
			t.Errorf("len(dbPsql.Jobs(%v, %v, %v, %v, %v, %v, %v)) = %v; want %v", product, version, attributes, name, measurement, timestamp, timestamp, got, want)
		}
		err = dbPsql.Delete(job)
		if err != nil {
			t.Errorf("error occurred: %+v", err)
		}
	})
	t.Run("", func(t *testing.T) {
		product := "p6"
		version := "v6"
		attributes := []string{"a61", "a62", "a63"}
		name := "n6"
		measurement := "m6"
		timestamp := time.Date(2017, 10, 29, 16, 18, 07, 0, time.UTC)
		value := 48
		createdOn := time.Now()
		createdBy := "192.168.1.2"
		job := &Job{product, version, attributes, name, measurement, timestamp, value}
		ctx := context.Background()
		ctx = WithCreatedOn(ctx, createdOn)
		ctx = WithCreatedBy(ctx, createdBy)
		gotErr := dbPsql.Create(ctx, job)
		wantErr := ErrJobAlreadyExists
		if errors.Cause(gotErr) != wantErr {
			t.Errorf("dbPsql.Jobs(%v, %v, %v, %v, %v, %v, %v) = %v; want %v", product, version, attributes, name, measurement, timestamp, timestamp, gotErr, wantErr)
		}
	})
}

func TestPsqlDelete(t *testing.T) {
	type args struct {
		job *Job
	}
	testsCases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				job: &Job{
					Product:     "pX",
					Version:     "v7",
					Attributes:  []string{"a71", "a72", "a74"},
					Name:        "n7",
					Measurement: "m7",
					Timestamp:   time.Date(2018, 2, 16, 2, 16, 51, 0, time.UTC),
					Value:       0,
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				job: &Job{
					Product:     "p7",
					Version:     "vX",
					Attributes:  []string{"a71", "a72", "a74"},
					Name:        "n7",
					Measurement: "m7",
					Timestamp:   time.Date(2018, 2, 16, 2, 16, 51, 0, time.UTC),
					Value:       0,
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				job: &Job{
					Product:     "p7",
					Version:     "v7",
					Attributes:  []string{"a7X", "a72", "a74"},
					Name:        "n7",
					Measurement: "m7",
					Timestamp:   time.Date(2018, 2, 16, 2, 16, 51, 0, time.UTC),
					Value:       0,
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				job: &Job{
					Product:     "p7",
					Version:     "v7",
					Attributes:  []string{"a71", "a72"},
					Name:        "n7",
					Measurement: "m7",
					Timestamp:   time.Date(2018, 2, 16, 2, 16, 51, 0, time.UTC),
					Value:       0,
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				job: &Job{
					Product:     "p7",
					Version:     "v7",
					Attributes:  []string{"a71", "a72", "a74"},
					Name:        "nX",
					Measurement: "m7",
					Timestamp:   time.Date(2018, 2, 16, 2, 16, 51, 0, time.UTC),
					Value:       0,
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				job: &Job{
					Product:     "p7",
					Version:     "v7",
					Attributes:  []string{"a71", "a72", "a74"},
					Name:        "n7",
					Measurement: "mX",
					Timestamp:   time.Date(2018, 2, 16, 2, 16, 51, 0, time.UTC),
					Value:       0,
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				job: &Job{
					Product:     "p7",
					Version:     "v7",
					Attributes:  []string{"a71", "a72", "a74"},
					Name:        "n7",
					Measurement: "m7",
					Timestamp:   time.Date(2020, 8, 4, 22, 50, 39, 0, time.UTC),
					Value:       0,
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				job: &Job{
					Product:     "p7",
					Version:     "v7",
					Attributes:  []string{"a71", "a72", "a74"},
					Name:        "n7",
					Measurement: "m7",
					Timestamp:   time.Date(2018, 2, 16, 2, 16, 51, 0, time.UTC),
					Value:       -1,
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				job: &Job{
					Product:     "p7",
					Version:     "v7",
					Attributes:  []string{"a71", "a72", "a74"},
					Name:        "n7",
					Measurement: "m7",
					Timestamp:   time.Date(2018, 2, 16, 2, 16, 51, 0, time.UTC),
					Value:       0,
				},
			},
			want: 0,
		},
	}
	product := "p7"
	version := "v7"
	attributes := []string{"a71", "a72", "a74"}
	name := "n7"
	measurement := "m7"
	timestamp := time.Date(2018, 2, 16, 2, 16, 51, 0, time.UTC)
	value := 0
	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			err := dbPsql.Delete(tt.args.job)
			if err != nil {
				t.Errorf("error occurred: %+v", err)
			}
			jobs, err := dbPsql.Jobs(product, version, attributes, name, measurement, timestamp, timestamp)
			if err != nil {
				t.Errorf("error occurred: %+v", err)
			}
			got := len(jobs)
			if got != tt.want {
				t.Errorf("dbPsql.Delete(%v) = %v; want %v", tt.args.job, got, tt.want)
			}
		})
	}
	createdOn := time.Now()
	createdBy := "192.168.1.2"
	job := &Job{product, version, attributes, name, measurement, timestamp, value}
	ctx := context.Background()
	ctx = WithCreatedOn(ctx, createdOn)
	ctx = WithCreatedBy(ctx, createdBy)
	err := dbPsql.Create(ctx, job)
	if err != nil {
		t.Errorf("error occurred: %+v", err)
	}
}

func TestPsqlProducts(t *testing.T) {
	got, err := dbPsql.Products()
	if err != nil {
		t.Errorf("error occurred: %+v", err)
	}
	want := []string{"p2", "p3", "p4", "p5", "p6", "p7", "pY"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("dbPsql.products() = %v; want %v", got, want)
	}
}

func TestPsqlVersion(t *testing.T) {
	type args struct {
		product string
	}
	testsCases := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				product: "p2",
			},
			want: []string{"v21", "v22"},
		},
		{
			name: "",
			args: args{
				product: "pX",
			},
			want: nil,
		},
	}
	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbPsql.Versions(tt.args.product)
			if err != nil {
				t.Errorf("error occurred: %+v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dbPsql.versions(%v) = %v; want %v", tt.args.product, got, tt.want)
			}
		})
	}
}

func TestPsqlAttributes(t *testing.T) {
	type args struct {
		product string
		version string
	}
	testsCases := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				product: "p3",
				version: "v3",
			},
			want: []string{"a31", "a32", "a33"},
		},
		{
			name: "",
			args: args{
				product: "pX",
				version: "v3",
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				product: "p3",
				version: "vX",
			},
			want: nil,
		},
	}
	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbPsql.Attributes(tt.args.product, tt.args.version)
			if err != nil {
				t.Errorf("error occurred: %+v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dbPsql.attributes(%v, %v) = %v; want %v", tt.args.product, tt.args.version, got, tt.want)
			}
		})
	}
}

func TestPsqlNames(t *testing.T) {
	type args struct {
		product    string
		version    string
		attributes []string
	}
	testsCases := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				product:    "p4",
				version:    "v4",
				attributes: []string{"a41", "a42"},
			},
			want: []string{"n41", "n42"},
		},
		{
			name: "",
			args: args{
				product:    "pX",
				version:    "v4",
				attributes: []string{"a41", "a42"},
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				product:    "p4",
				version:    "vX",
				attributes: []string{"a41", "a42"},
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				product:    "p4",
				version:    "v4",
				attributes: []string{"a4X", "a42"},
			},
			want: nil,
		},
	}
	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbPsql.Names(tt.args.product, tt.args.version, tt.args.attributes)
			if err != nil {
				t.Errorf("error occurred: %+v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dbPsql.names(%v, %v, %v) = %v; want %v", tt.args.product, tt.args.version, tt.args.attributes, got, tt.want)
			}
		})
	}
}

func TestPsqlMeasurements(t *testing.T) {
	type args struct {
		product    string
		version    string
		attributes []string
		name       string
	}
	testsCases := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				product:    "p5",
				version:    "v5",
				attributes: []string{"a51", "a52"},
				name:       "n5",
			},
			want: []string{"m51", "m52"},
		},
		{
			name: "",
			args: args{
				product:    "pX",
				version:    "v5",
				attributes: []string{"a51", "a52"},
				name:       "n5",
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				product:    "p5",
				version:    "vX",
				attributes: []string{"a51", "a52"},
				name:       "n5",
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				product:    "p5",
				version:    "v5",
				attributes: []string{"a5X", "a52"},
				name:       "n5",
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				product:    "p5",
				version:    "v5",
				attributes: []string{"a51", "a52"},
				name:       "nX",
			},
			want: nil,
		},
	}
	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbPsql.Measurements(tt.args.product, tt.args.version, tt.args.attributes, tt.args.name)
			if err != nil {
				t.Errorf("error occurred: %+v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dbPsql.measurements(%v, %v, %v, %v) = %v; want %v", tt.args.product, tt.args.version, tt.args.attributes, tt.args.name, got, tt.want)
			}
		})
	}
}

func TestPsqlJobs(t *testing.T) {
	type args struct {
		product     string
		version     string
		attributes  []string
		name        string
		measurement string
		lbound      time.Time
		rbound      time.Time
	}
	testsCases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				product:     "p6",
				version:     "v6",
				attributes:  []string{"a61", "a62"},
				name:        "n6",
				measurement: "m6",
				lbound:      time.Date(2017, 8, 19, 8, 40, 23, 0, time.UTC),
				rbound:      time.Date(2018, 7, 2, 6, 34, 54, 0, time.UTC),
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				product:     "p6",
				version:     "v6",
				attributes:  []string{"a61", "a62"},
				name:        "n6",
				measurement: "m6",
				lbound:      time.Date(2017, 10, 5, 12, 57, 44, 0, time.UTC),
				rbound:      time.Date(2017, 10, 5, 12, 57, 44, 0, time.UTC),
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				product:     "p6",
				version:     "v6",
				attributes:  []string{"a61", "a62"},
				name:        "n6",
				measurement: "m6",
				lbound:      time.Date(2019, 3, 31, 11, 12, 16, 0, time.UTC),
				rbound:      time.Date(2019, 4, 6, 14, 20, 19, 0, time.UTC),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				product:     "pX",
				version:     "v6",
				attributes:  []string{"a61", "a62"},
				name:        "n6",
				measurement: "m6",
				lbound:      time.Date(2017, 8, 19, 8, 40, 23, 0, time.UTC),
				rbound:      time.Date(2018, 7, 2, 6, 34, 54, 0, time.UTC),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				product:     "p6",
				version:     "vX",
				attributes:  []string{"a61", "a62"},
				name:        "n6",
				measurement: "m6",
				lbound:      time.Date(2017, 8, 19, 8, 40, 23, 0, time.UTC),
				rbound:      time.Date(2018, 7, 2, 6, 34, 54, 0, time.UTC),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				product:     "p6",
				version:     "v6",
				attributes:  []string{"a6X", "a62"},
				name:        "n6",
				measurement: "m6",
				lbound:      time.Date(2017, 8, 19, 8, 40, 23, 0, time.UTC),
				rbound:      time.Date(2018, 7, 2, 6, 34, 54, 0, time.UTC),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				product:     "p6",
				version:     "v6",
				attributes:  []string{"a61", "a62"},
				name:        "nX",
				measurement: "m6",
				lbound:      time.Date(2017, 8, 19, 8, 40, 23, 0, time.UTC),
				rbound:      time.Date(2018, 7, 2, 6, 34, 54, 0, time.UTC),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				product:     "p6",
				version:     "v6",
				attributes:  []string{"a61", "a62"},
				name:        "n6",
				measurement: "mX",
				lbound:      time.Date(2017, 8, 19, 8, 40, 23, 0, time.UTC),
				rbound:      time.Date(2018, 7, 2, 6, 34, 54, 0, time.UTC),
			},
			want: 0,
		},
	}
	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			jobs, err := dbPsql.Jobs(tt.args.product, tt.args.version, tt.args.attributes, tt.args.name, tt.args.measurement, tt.args.lbound, tt.args.rbound)
			if err != nil {
				t.Errorf("error occurred: %+v", err)
			}
			got := len(jobs)
			if got != tt.want {
				t.Errorf("len(dbPsql.Jobs(%v, %v, %v, %v, %v, %v, %v)) = %v; want %v", tt.args.product, tt.args.version, tt.args.attributes, tt.args.name, tt.args.measurement, tt.args.lbound, tt.args.rbound, got, tt.want)
			}
		})
	}
}
