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
# Story?

<!-- Notes
Erik
It would be nice to have a story. Any ideas?
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
- Four slides, one per metric.
- For each metric show a demo and/or example CDEvents relevant for the metric and/or diagram.
- Spend some time on the data in events, discuss how it can be used to correlate events.
Andrea & Erik? We could split, some metrics each
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

-->

---

# Lead Time for Changes

<!-- Notes

Examples:
- Kaniko, Buildah
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
- the change ID (change event)
- the latest change ID (build event)

(TBD)

-->

# Change Failure Rate

---

# Time to Restore Service

---

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
