package routing_table_test

import (
	. "github.com/cloudfoundry-incubator/route-emitter/routing_table"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MessagesToEmit", func() {
	var (
		messagesToEmit MessagesToEmit
		messages1      []RegistryMessage
	)

	BeforeEach(func() {
		messagesToEmit = MessagesToEmit{}
		messages1 = []RegistryMessage{
			{
				Host: "1.1.1.1",
				Port: 61000,
				App:  "log-guid-2",
				URIs: []string{"host1.example.com"},
			},
			{
				Host: "1.1.1.1",
				Port: 61001,
				App:  "log-guid-1",
				URIs: []string{"host1.example.com"},
			},
			{
				Host: "1.1.1.1",
				Port: 61003,
				App:  "log-guid-2",
				URIs: []string{"host2.example.com", "host3.example.com"},
			},
			{
				Host: "1.1.1.1",
				Port: 61004,
				App:  "log-guid-3",
				URIs: []string{"host3.example.com"},
			},
		}
	})

	Describe("RouteRegistrationCount", func() {
		Context("when there are registration messages", func() {
			BeforeEach(func() {
				messagesToEmit.RegistrationMessages = messages1
			})

			It("adds the number of hostnames in each route message", func() {
				Ω(messagesToEmit.RouteRegistrationCount()).Should(BeEquivalentTo(5))
			})
		})

		Context("when registration messages is nil", func() {
			BeforeEach(func() {
				messagesToEmit.RegistrationMessages = nil
				messagesToEmit.UnregistrationMessages = messages1
			})

			It("adds the number of hostnames in each route message", func() {
				Ω(messagesToEmit.RouteRegistrationCount()).Should(BeEquivalentTo(0))
			})
		})
	})

	Describe("RouteUnregistrationCount", func() {
		Context("when there are unregistration messages", func() {
			BeforeEach(func() {
				messagesToEmit.UnregistrationMessages = messages1
			})

			It("adds the number of hostnames in each route message", func() {
				Ω(messagesToEmit.RouteUnregistrationCount()).Should(BeEquivalentTo(5))
			})
		})

		Context("when registration messages is nil", func() {
			BeforeEach(func() {
				messagesToEmit.RegistrationMessages = messages1
				messagesToEmit.UnregistrationMessages = nil
			})

			It("adds the number of hostnames in each route message", func() {
				Ω(messagesToEmit.RouteUnregistrationCount()).Should(BeEquivalentTo(0))
			})
		})
	})
})
