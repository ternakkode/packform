package seed

import (
	"context"
	"log"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/ternakkode/packform-backend/internal/config"
	"github.com/ternakkode/packform-backend/pkg/bunclient"
)

type Seed struct{}

func Exec(seedNames []string) {
	config.Init()
	bunclient.InitDB(&config.GetConfig().DB)

	s := &Seed{}
	wg := &sync.WaitGroup{}

	if len(seedNames) > 0 {
		for _, name := range seedNames {
			if strings.HasSuffix(name, "Seeder") {
				wg.Add(1)
				go func(n string, wg *sync.WaitGroup) {
					ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

					defer func() {
						// if r := recover(); r != nil {
						// 	log.Printf("failed to run seeder func %s, %v", n, r)
						// }

						wg.Done()
						cancel()
					}()

					meth := reflect.ValueOf(s).MethodByName(n)
					if !meth.IsValid() {
						log.Printf("failed to run seeder func %s, invalid function", n)
						return
					}

					err := meth.Call([]reflect.Value{reflect.ValueOf(ctx)})
					if !err[0].IsNil() {
						log.Printf("failed to run seeder func %s, %v", n, err[0])
						return
					}

					log.Printf("successfully run seeder func %s", n)
				}(name, wg)

			}
		}
	}

	wg.Wait()
}
