---
marp: true
theme: gaia
backgroundColor: #112
class: invert
_footer: 'cdCon | 07.06.2022'
---
<!-- Uses MARP, see https://marp.app/ -->

<!--
class:
 - lead
 - invert
-->

# cdCon 2022

#### Building DevOps metrics for your choice of CD tools through CDEvent

---
# A long long time ago

<!-- Notes
Erik

In October of 2012 I was working in the telecom industry and I had just gotten
involved in one of the earlier CI/CD efforts in my workplace.

I got my manager to explain the current situation and he said, "Well, Erik, to
start with we make releases every six months, right? The development of those
releases typically start around a year before release, and the releases get
verified by the verification teams during the last three months before release."
-->

---
# A long long time...

<!-- Notes

So, this means developers frequently need to wait a few months for their work to
be fully tested. And, of course, it is unlikely that whatever they are working
on right now is directly related to what is being verified by the verification
team. If a bug is discovered, developers need to make the mother of all
context switches to find and fix it. Not at all good.

My manager then proclaimed "We need to bring this time, the time from
development to verification and release, down from a few months to a few hours."

I remember thinking "going from months to hours seems like a crazy big step!"
but I have since been shown over and over again that this step is actually not
that crazy!

One factor that really helps taking this step is to understand where our
bottlenecks are, and for that we need metrics.

My name is Erik Sternerson...
-->

---
# cdCon 2022

#### Building DevOps metrics for your choice of CD tools through CDEvent

####

####

#### Andrea Frittoli, IBM. Erik Sternerson, doWhile

---
<!--
_class:
 - invert
-->

# In this talk
<!-- Comment
-->

<!-- Notes
Andrea
-->

* ## Learn About CDEvents

* ## Learn About DevOps Metrics

* ## How do they fit together

---

# CDEvents

<!-- TODO
A couple of slides
- Project goals (and history?)
- Use cases
-->

<!-- Notes
Erik
-->

---

# Conceptual: Common language

<!-- The conceptual goal of the CDEvents project is to help build a common
language for CI/CD and surrounding domains.

-->

---

<!--
_class:
 - invert
-->

# Wider effort

<!-- So I set "help build" just now, and that is because this it is not only
CDEvents involved in this work. -->

* ## CDF SIG Interoperability

<!-- The Interoperability special interests group of the Continuous Delivery
foundation can probably be seen as the "driver" of establishing this common
language. 

This group does a lot of work defining and establishing terms for
similar concepts across the CI/CD ecosystem, and a lot of these terms pop up in
the CDEvents project in one way or another.
 -->

* ## CDF SIG Events

<!-- The Events special interest group spawned out of the interoperability
group in {TODO DATE late 2020} as a workgroup focusing specifically on a
vocabulary for events in CI/CD. 

It became a full SIG {TODO DATE about a year later}, and is the root of...
 -->

---

# CDEvents

<!-- The CDEvents project, and its concrete goal: -->

---

# Concrete: Spec and SDKs

<!-- To build a specification for events in CI/CD, and to build a set of SDKs
that help others send and receive such events.
-->

---

# The spec

<!-- Lets dig in to the spec a bit first. -->

<!-- TODO Add a circular diagram where the spec feeds into the SDKs,
the SDKs feed into the integrations and PoCs and they feed back to the spec.
Show this step by step for spec, SDKs and PoCs. -->

---

<!-- TOOD: Add multi-image diagram showing how the spec declares multiple event 
types, their event data and how it uses the CloudEvents base -->

---

# The SDKs

<!-- Given this spec, we can now work on a set of SDKs for multiple
programming languages and platforms. -->

---

<!-- TOOD: Add multi-image code example from Go SDK -->

---

# The integrations and PoCs

<!-- Finally, with the SDKs, we can wokr on integratins CDEvents into 
new and existing tools and solutions, and set up various proof-of-concepts
to test out new ideas and help drive the specification forward. -->

---

<!-- TOOD: Add multi-image code example from Go SDK -->

---

<!--
_class:
 - invert
-->

# Wider effort (again)

<!-- And this AGAIN is a wider effort. -->

* ## CDEvents team

<!-- The spec itself is driven by the CDEvents team, but with plenty of
support, input and feedback from the wider community. -->

* ## Project communities

<!-- Several integrations and proof-of-concepts thus far have been done 
wholly or partly outside of the CDEvents project itself, by members of the
communit for the projects  -->

---

# DevOps (DORA) Metrics

<!-- TODO
A couple of slides
- Intro to DevOps metrics, high performing teams, measuring performance
- Specifically, DORA metrics
-->

<!-- Notes
Andrea
-->

---

# Metrics through CDEvents

<!--TODO
-->

<!-- Notes
- 4+ slides, one or more per metric.
- For each metric show a demo and/or example CDEvents relevant for the metric and/or diagram.
- Spend some time on the data in events, discuss how it can be used to correlate events.

Andrea
-->

---

# Deployment Frequency

<!-- Notes

Let's take a few examples:
- kubectl or another deployment tool in a Tekton Task
- gitops tool like ArgoCD and Flux
- spinnaker
- keptn

The relevant data here is:
- environment (deployment event)
- artifact name (deployment event)

Erik
-->

---

# Lead Time for Changes

<!-- Notes

Examples:
- Kaniko, Buildah for container images
- Tekton, Jenkins, Shipwright

Assuming a single artifact, single branch, how this metric is
calculated still depends on the versioning scheme used for the
artifact. No back-porting means that a change is always included
in the next build after to the change is merged, and in all builds
after that. If the build model is more complex, we must rely on the
change ID, the latest change ID from the build, and ask the SCM if
the change ID was merged before the build change ID.

The relevant data is:
- the timestamp (change and build events)
- the repository (change and build events)
- the latest change ID (build event)
- the change ID (change event)

Deployment tools, that Erik introduced, take a specific build and
deploy it to production. The artifact name is not enough, we need
the artifact ID so that we may associate a specific artifact and thus
specific changes.

The relevant data is:
- the artifact ID (deploy and build events)

In real life, we will often need to consider composition scenarios,
where an artifact is not directly deployed, but it's used instead to
build a composite artifact or collection of artifacts (release).

We started investigating how to define such scenarios in CDEvents,
exploring the idea of composition.

Andrea
-->

# Change Failure Rate

<!--

Erik
-->
---

# Time to Restore Service

<!--


Andrea
-->

---
# Key takeaways

---

# TBW
<!-- Notes

Something along the lines of:

- Metrics are hard regardless
- A common language can help
- Call to action
-->

---

# TBW

---

# TBW

---

# Thank you!

---
<!--
_footer: 'cdCon | 07.06.2022'
-->

# Questions?

##
##
##
##
##

#### Andrea Frittoli, IBM. Erik Sternerson, doWhile
