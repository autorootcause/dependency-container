# dependency-container
TLDR: Easy-to-install sidecar to validate service dependencies

All too often, debugging gnarly microservices issues involves several steps:
1. get production approval
2. ssh into the bastion  
3. ssh into the container
4. run curl command or script (hopefully you don't misspell the command in the middle of the night)

You might want to test if you can `ping a dependency endpoint` from that individual host. Or you want to `renew a certificate` or `retokenike the pupper key`.

Runbooks are filled with these commands, but I find the process of following runbooks cumbersome, error-prone and frankly annoying when I get paged at 4 in the morning. Sometimes, the runbooks are wrong, missing information, or simply not updated and I simply don't use them as a source of truth. Fixing this is a culture/process problem and I humbly think stocks vest faster than culture changes.

However, I think container logs are a great source of truth though. I want the container to run validating scripts with my dependencies, check the expiry of the certificate, etc. before I even get paged.
This way, when I actually get paged, I'm not sitting around waiting for an approval or parsing through logs/metrics/traces in the middle of night and I can figure out the source of the issue faster.

In that spirit, I offer an alternative: why not have a pre-configured set of scripts that run when the readiness/liveness/startup probes break and output the responses to your preferred logging sink, e.g. Splunk.

Ultimately, debugging is about validating hypotheses and these commands are effectively testing some aspect of the system to narrow down the issue.

For example, running a command to test if you can connect to AWS tells you:

- cluster is still working
- network policy for outbound requests are set correctly
- NAT gateway/outbound proxy is working normally
- AWS is not down
- more stuff in between

As Pranav Mistry says: `we as humans are not interested in technology [or processes], we're interested in information.`
In that mode of thinking, I have a shameless plug: reach out to us if you'd like to go a step forward to use these hypotheses to automagically root cause incidents on your phone.

# Design

This library is designed with a plugin system in mind. A plugin is an abstraction that can be:

- pre-built plugin, e.g. `AWS-STS`
- shell command, e.g. `curl -X autorootcause.com`
- script, e.g. `sh renew_certificate.sh`

### Pre-built Plugins

With the number of folks building managed services, there's probably a lot of duplication in how folks validate whether a service is up. So if you'd like to use one of the following pre-built plugins, add one of the plugins to your config, as follows:
```yaml
plugins: [
  name: AWS-STS,
  type: prebuilt,
  ExpectedResult: http_2xx,
  trigger:
    trigger_type: [
      on_liveness_failure,
      periodic,
    ],
    trigger_interval: 5,
    trigger_timeout: 60s
]
```

### Shell Command

Shell commands 


