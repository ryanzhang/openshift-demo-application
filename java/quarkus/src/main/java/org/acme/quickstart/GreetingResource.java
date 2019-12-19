package org.acme.quickstart;

import javax.inject.Inject;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

import org.eclipse.microprofile.config.inject.ConfigProperty;
import org.jboss.resteasy.annotations.jaxrs.PathParam;

@Path("/")
public class GreetingResource {

    @Inject
    GreetingService service;

    @ConfigProperty(name = "MESSAGE_PREFIX", defaultValue = "Hi")
    String messagePrefix;
    
    @GET
    @Produces(MediaType.TEXT_PLAIN)
    @Path("/greeting/{name}")
    public String greeting(@PathParam String name) {
        return service.greeting(name);
    }

    @GET
    @Path("/")
    @Produces(MediaType.TEXT_HTML)
    public String hello() {
        return "<marquee bgcolor='red'><h1><pre>" + messagePrefix;
    }
}