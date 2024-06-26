package main

import (
	"fmt"

	"github.com/murshid/Interface"
	"github.com/murshid/Range"
	"github.com/murshid/Select"
	"github.com/murshid/Switch"
	T "github.com/murshid/Tickers"
	"github.com/murshid/Timeouts"
	"github.com/murshid/array"
	"github.com/murshid/atomic_counters"
	"github.com/murshid/buffered_channel"
	"github.com/murshid/channel_direction"
	"github.com/murshid/channel_sync"
	"github.com/murshid/channels"
	"github.com/murshid/close_channel"
	closure "github.com/murshid/closures"
	"github.com/murshid/constants"
	"github.com/murshid/custom_error"
	"github.com/murshid/enums"
	"github.com/murshid/errors"
	"github.com/murshid/for_loop"
	"github.com/murshid/function"
	"github.com/murshid/generics"
	go_routine "github.com/murshid/goroutines"
	hello_world "github.com/murshid/hello_world"
	"github.com/murshid/if_else"
	"github.com/murshid/maps"
	"github.com/murshid/methods"
	"github.com/murshid/multiple_return"
	"github.com/murshid/mutex"
	"github.com/murshid/non_Blocking_channels"
	"github.com/murshid/pointers"
	"github.com/murshid/range_on_channel"
	"github.com/murshid/rate_limiter"
	"github.com/murshid/recursion"
	"github.com/murshid/slices"
	"github.com/murshid/stringsnrune"
	"github.com/murshid/struct_embedding"
	"github.com/murshid/structs"
	"github.com/murshid/timer"
	value "github.com/murshid/values"
	"github.com/murshid/variables"
	"github.com/murshid/variadic_function"
	Waitgroups "github.com/murshid/waitgroups"
	Worker_pools "github.com/murshid/worker_pools"
)

func main() {
	hello_world.Hello("hello world")
	value.Values()
	T.Ticker()
	array.Array()
	atomic_counters.Atomic_counters()
	buffered_channel.Buffered_channel()
	channel_direction.Channel_direction()
	channel_sync.Channel_sync()
	channels.Channels()
	close_channel.Close_channel()
	closure.Closure()
	constants.Constants()
	custom_error.Custom_error()
	enums.Enums()
	errors.Errors()
	for_loop.For_loop()
	function.Function()
	generics.Generics()
	go_routine.Go_routine()
	fmt.Println(if_else.If_else(5))
	Interface.Interface()
	maps.Maps()
	methods.Methods()
	multiple_return.Multiple_return()
	mutex.Mutex()
	non_Blocking_channels.Non_Blocking_Channel()
	pointers.Pointers()
	Range.Range()
	range_on_channel.Range_on_channel()
	rate_limiter.Rate_limiter()
	recursion.Recursion()
	Select.Select()
	slices.Slices()
	stringsnrune.Stringsnrune()
	struct_embedding.Embedding()
	structs.Structs()
	fmt.Println(Switch.Switch(13))
	Timeouts.Timeout()
	timer.Timer()
	variables.Variables()
	variadic_function.Var_func()
	Waitgroups.Waitgroup()
	Worker_pools.Worker_pool()
}
