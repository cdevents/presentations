---
marp: true
theme: gaia
backgroundColor: #112
class: invert
---

<!-- Uses MARP, see https://marp.app/ -->

<!--
class: 
 - lead
 - invert
-->

# CDEvents

### Interoperability in the Continuous Delivery space

---

<!-- Erik -->

![bg contain](images/glue_code.jpg)

<!-- Glue code story:
Back in November 2019 I was asked by a colleague to help prepare a detailed training session on our continuous delivery setup. *This was for the autonomous drive development at Volvo Cars, so quite a complex project.*

By then we had been working on this continous delivery setup for about a year and a half and a simplified overview of what we had **built** was needed. After an hour in front of this massive whiteboard drawing boxes for everything we were actually a bit surprised to realize that we had over 70 components in this setup. 

Some of these components were configuration files, some were existing solutions like Jenkins, Gerrit and Nexus, but about **25** of the components we identified were basically **glue code**. Like moving information from one component to an other, or allowing one component to run and get results from an other. 

I remember thinking "this work is probably repeated in other projects and organizations too, **there has to be a better way**. And as it turns out, **there is**!
-->

---

<!-- Andrea -->

# CDEvents

### Interoperability in the Continuous Delivery space

<br />
<br />

##### Andrea Frittoli & Erik Sternerson

<!-- Today we will talk about CDevents, a project to provide a protocol and vocabulary for events and signalling in the Continuous Delivery area. -->

---

<!--
class: 
 - invert
-->

# After this talk
<!--  -->

* ## Know the basics of Continuous Delivery events
<!--  -->
* ## Understand use cases for events
<!--  -->
* ## Know about the CDEvents project
<!--  -->
---

<!--
_class: 
 - lead
 - invert
-->

# Continuous Delivery Events

<!-- Erik -->

<!-- First up, we'll look into Continous Delivery Events as a concept.-->

---

# A primer on Continuous Delivery

<!-- Let's start by defining Continuous Delivery -->

* ### Ready-for-release
<!-- The same way that Continouos Integration aims to get your code integrated to the main branch, Continuous Delivery aims to get your project in a state where it can be delivered -->
* ### From idea to metrics and back again
<!-- In its widest form, Continuous Delivery supports you all the way from breaking down your initial idea through getting metrics from your release project, and using those metrics as the basis for your next idea. -->
* ### Too big for any one system to solve
<!-- With activities ranging from verification and validation through packaging, release and remediation, it is not feasible to have a single "one-size-fits-all" Continouos Delivery system that solves everything. Instead, we need to rely on integrations. -->

---

![bg contain](images/Continuous-Delivery-Landscape.png)

<!-- Here are some examples of open-source projects in the continuous delivery area, many of which will at some point need to be glued together for one purpose or another. -->

---

# Continuous Delivery Events

* ### When something valuable happens, __announce__ it
<!-- Value is produced in Continuous Delivery all the time. And we are not only talking binaries and artifacts here, but also actions being taken and activities starting and stopping. -->
* *Change, Build, Test, Deployment, New Env., Failure, Remediation, etc. etc.*
* ### When something interesting happens, __react__ to it
<!-- These events allow any system listening to these events to take arbitrary action when something they are interested in happens. -->
* *Visualize, Store, Announce, Run, Retry, etc. etc.*

---

# Not a brand new idea

- ## Argo Events

- ## Eiffel

- ## Keptn

---

<!--
class: 
 - lead
 - invert
-->

# "Let's work together!"

---

<!-- Andrea -->

# [CDEvents Logo]

---

<!--
class: 
 - invert
-->

# What is CDEvents?
<!-- TBW -->

* ## A vocabulary
<!-- TBW -->

* ## A specification
<!-- TBW -->

* ## A project
<!-- TBW -->

* ## In progress :)
<!-- TBW -->

---

# The CDEvents project

* ## SIG-Interop @ CDF mid 2020
<!-- TBW -->

* ## SIG-Events @ CDF early 2021
<!-- TBW -->

* ## CDEvents late 2021
<!-- TBW -->

* ## <https://github.com/cdevents>

---

# The CDEvents vocabulary

[Insert vocabulary-relevant picture here]

<!-- TBW -->

---

# The CDEvents specification

[Insert JSON event example picture]

<!-- TBW -->

---

<!--
class: 
 - lead
 - invert
-->

# CDEvents use cases

---

# Interoperability

<!-- Andrea -->

---

<!--
class: 
 - invert
-->

# TBW Use Case 1 Details

TBW

---

<!--
_class: 
 - lead
 - invert
-->

# Observability & Metrics

<!-- Erik -->

---

# TBW Use Case 2 Details

TBW

---

<!--
class: 
 - lead
 - invert
-->

<!-- Andrea -->

# What can CDEvents do for you?

---

# TBW Do for you 1

---

# TBW Do for you 2

---

# TBW Do for you 3

---

# The way forward

---

<!--
class: 
 - invert
-->

# CDEvents way forward

* ## Learn more about what is used
<!-- TBW -->

* ## Keep building the spec and the SDKs
<!-- TBW -->

* ## Implement into existing tools/services
<!-- TBW -->

---

<!--
class: 
 - lead
 - invert
-->

<!-- Erik -->

# Key Takeaways

---

# Spec for interoperability and observability

---

# TBW Key Takeaway 2

---

# TBW Key Takeaway 3

---

# Thank You!

---

# Questions?

<br />
<br />
<br />

### _Join the cause!_
### _CDEvents <https://github.com/cdevents>_