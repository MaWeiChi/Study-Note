package main

go func() {
	for !ntp.cancel {
		ctx, cancel := context.WithCancel(context.Background())
		go ntp.Update(ctx)
		go func{
			for {
				select {
					case <-ctx.Done():
						break
					case <-time.After(5):
						if ntp.cancel{
							ntp.wg.Done()
							cancel()
							return
						}
						
				}
			}
	}
	select{
		case <-time.After(60 * time.Millisecond):
			cancel()
	}
		// for i := 0; i < ntp.Interval; i += 5 {
		// 	if ntp.cancel {
		// 		ntp.wg.Done()
		// 		cancel()
		// 		return
		// 	}
		// 	time.Sleep(time.Duration(5) * time.Second)
		// }
		cancel()
	}
	ntp.wg.Done()
}()
}