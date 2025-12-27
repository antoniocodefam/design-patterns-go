package main

type motorcycleFactory struct{
	subscribers map[string][]ISubscriber
}

func newMotorcycleFactory() *motorcycleFactory {
	return &motorcycleFactory{
		subscribers: map[string][]ISubscriber{},
	}
}

func (f *motorcycleFactory) createMotorcycle(brand, model string) {
	subscribers, ok := f.subscribers[brand]

	if ok {
		for _, subscriber := range subscribers {
			subscriber.notify(brand, model)
		}
	}
}

func (f *motorcycleFactory) unsubscribe(subscriber ISubscriber, brand string) {
	subscribers, ok := f.subscribers[brand] 

	if ok {
		newSubscribers := make([]ISubscriber, 0)
		for _, s := range subscribers {
			if s.getName() != subscriber.getName() {
				newSubscribers = append(newSubscribers, s)
			}
		}

		f.subscribers[brand] = newSubscribers
	}
}

func (f *motorcycleFactory) subscribe(subscriber ISubscriber, brand string) {
	f.subscribers[brand] = append(f.subscribers[brand], subscriber)
}