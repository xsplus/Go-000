package main

import (
    "context"
    "fmt"
    "golang.org/x/sync/errgroup"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

func main()  {
    g, ctx := errgroup.WithContext(context.Background())
    g.Go(func() error {
        return startServer(ctx)
    })
    g.Go(func() error {
        return listenSignal(ctx)
    })
    if err := g.Wait(); err != nil {
        fmt.Println("errgroup return error : ", err)
    }
    fmt.Println("EXIT!")
}

func listenSignal(ctx context.Context) error  {
    c := make(chan os.Signal)
    signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
    log.Println("listen signal ...")
    select {
    case s := <-c:
        return fmt.Errorf("get %v signal", s)
    case <-ctx.Done():
        return fmt.Errorf("signal：other work done")
    }
}

func startServer(ctx context.Context) error {
    server := &http.Server{Addr:":8081", Handler: http.DefaultServeMux}

    go func() {
        <-ctx.Done()
        log.Println("http: other work done")
        e := server.Shutdown(ctx)
        log.Println(e)
    }()

    log.Println("start server")
    return server.ListenAndServe()
}