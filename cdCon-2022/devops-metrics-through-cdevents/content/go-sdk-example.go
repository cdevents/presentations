		// Create an event.
		event, _ := cde.CreateServiceEvent(
			cde.ServiceDeployedEventV1, 
			serviceEnvId,
			serviceName, 
			serviceVersion, 
			serviceData)

		// Set a source.
		event.SetSource(source)

		// Set a target.
		ctx := cloudevents.ContextWithTarget(
			context.Background(), CDE_SINK)