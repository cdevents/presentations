---
marp: true
theme: gaia
backgroundColor: #112
class: invert
---
<!-- Erik -->

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

By then we had been working on this continuous delivery setup for about a year and a half and a simplified overview of what we had **built** was needed. After an hour in front of this massive whiteboard drawing boxes for everything we were actually a bit surprised to realize that we had over 70 components in this setup.

Some of these components were configuration files, some were existing solutions like Jenkins, Gerrit and Nexus, but about **25** of the components we identified were basically **glue code**. Like moving information from one component to an other, or allowing one component to run and get results from an other.

I remember thinking "this work is probably repeated in other projects and organizations too, **there has to be a better way**. And as it turns out, **there is**!
-->

---
<!-- Andrea + Erik -->

# CDEvents

### Interoperability in the Continuous Delivery space

<br />
<br />

##### Andrea Frittoli & Erik Sternerson

<!-- Thank you Erik, I loved that story! Indeed we believe there is a better way, and that's what we're going to present today: we will talk about CDevents, a project to provide a protocol and vocabulary for events and signalling in the Continuous Delivery area.
Before we dive in, let's briefly introduce ourselves.

Hi, my name is Andrea Frittoli, I'm a software engineer at IBM. I serve as co-chair for the Events special interest group at the at the continuos delivery foundation and I'm one of the maintainer of the CDEvents specification. I'm also a maintainer of Tekton.

[Erik]
-->

---
<!-- Andrea -->
<!--
class:
 - invert
-->

# After this talk
<!--  -->

* ## Know the basics of Continuous Delivery events
* ## Understand use cases for events
* ## Know about the CDEvents project
* ## Share your Glue-code story

<!-- Today we're in the FOSDEM CI/CD devroom, so most of you are probably familiar to some degree with the concept of continuous delivery.
Today we we'll focus on events in the continuous delivery space and how do we use them to solve the glue-code issue, and more. We'll also introduce the CDEvents project, specification and community.
We have a few use cases lined up to present, and we'd love to hear your story as well. Do you use multiple platforms in your CD setup? Which ones? Do they talk to each other and how? Use the link from the slides and in the chat to tell us your story.
We'd love also for you to join one of our working groups and discuss your use cases or how you use events in your setup.
But let's get started - Erik -->
---
<!-- Erik -->
<!--
_class:
 - lead
 - invert
-->

# Continuous Delivery Events

<!-- First up, we'll look into Continuous Delivery Events as a concept.-->

---
<!-- Erik -->
# A primer on Continuous Delivery

<!-- Just to make sure everyone is on the same page, let's start by defining Continuous Delivery -->

* ### Ready-for-release
<!-- The same way that Continuous Integration aims to get your code integrated to the main branch, Continuous Delivery aims to get your project in a state where it can be delivered -->
* ### From idea to metrics and back again
<!-- In its widest form, Continuous Delivery supports you all the way from breaking down your initial idea through getting metrics from your release project, and using those metrics as the basis for your next idea. -->
* ### Too big for any one system to solve
<!-- With activities ranging from verification and validation through packaging, release and remediation, it is not feasible to have a single "one-size-fits-all" Continuous Delivery system that solves everything. Instead, we need to rely on integrations. -->

---
<!-- Erik -->
![bg contain](images/Continuous-Delivery-Landscape.png)

<!-- Here are some examples of open-source projects in the continuous delivery area, many of which will at some point need to be glued together for one purpose or another. -->

---
<!-- Erik -->
# Continuous Delivery Events

* ### When something valuable happens, __announce__ it
<!-- Value is produced in Continuous Delivery all the time. And we are not only talking binaries and artifacts here, but also actions being taken and activities starting and stopping. -->
* *Change, Build, Test, Deployment, New Env., Failure, Remediation, etc. etc.*
* ### When something interesting happens, __react__ to it
<!-- These events allow any system listening to these events to take arbitrary action when something they are interested in happens. -->
* *Visualize, Store, Announce, Run, Retry, etc. etc.*

---
<!-- Erik -->
# Not a brand new idea

- ## Argo Events

- ## Eiffel

- ## Keptn

---
<!-- Erik -->
<!--
class:
 - lead
 - invert
-->

# "Let's work together!"

---
<!-- Andrea -->
<!--
backgroundColor: #FFF
-->

![cdevents logo](3rdPartyLogos/cdevents-draft.png)

<!-- Enter the CDEvents project.
Now that we're recording, the logo is almost finalized. We'll update the slides with the final version of the logo as soon as it's available.-->

---
<!-- Andrea -->

<!--
backgroundColor: #112
class:
 - invert
-->

# What is CDEvents?

* ## A vocabulary
* ## A specification
* ## A project
* ## In progress :)
* ## [github.com/cdevents](https://github.com/cdevents)

<!-- At the time we are recording, CDEvents is the latest addition to the CDF incubating projects. CDEvents aims to be the common language the CD ecosystem events, so it provides a vocabulary, a specification as well as SDKs. It's work in progress but already a few projects provide experimental support for it.-->
---
<!-- Andrea -->

# The CDEvents project

* ## SIG-Interop @ CDF mid 2020
* ## SIG-Events @ CDF early 2021
* ## CDEvents late 2021

<!-- Before we dive in further, a bit of history. The interoperability special interest group was created in mid 2020 with the goal to discuss and research interoperability in the CD space. One of the work-stream of the SIG was focused on interoperability through events. In early 2021 we transformed the work-stream into an SIG of its own, and towards the end of last year we created the CDEvents project, to host the specification work of the SIG. The project was proposed as a CDF incubating project and accepted in December 2021. -->
---
<!-- Andrea -->

# The CDEvents vocabulary

* ## Core
* ## Source Code Version Control
* ## Continuous Integration Pipelines
* ## Continuous Deployment Pipelines
* ## ...more to come

<!-- Most CD platform define their abstractions, data model and nomenclature. The interoperability SIG has already been collecting these from various platforms. Many names are shared across platforms, but sometimes the same name bears different meanings in different projects. To achieve interoperability through events we needed to define a nomenclature with shared semantics across platforms. To achieve that we first created four buckets to group the different events. The list is not exhaustive, as we'd like eventually to cover more aspects of the software lifecycle, like for instance monitoring and incident management. Within each bucket, we defined a few abstractions. For instance "Core" defines "Task Runs" and "Pipeline Runs". Continuous integration defines "Build", "Test Case", "Test Suite" and "Artifact". -->

---
<!-- Andrea -->

# The CDEvents specification

[Insert JSON event example picture]

<!-- TBW -->

---
<!-- Andrea -->

<!--
class:
 - lead
 - invert
-->

# CDEvents use cases

---
<!-- Erik -->

# Interoperability

---
<!-- Erik -->

<!--
class:
 - invert
-->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_1.svg)

<!-- TBW -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_2.svg)

<!-- TBW -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_3.svg)

<!-- TBW -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_4.svg)

<!-- TBW -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_5.svg)

<!-- TBW -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_6.svg)

<!-- TBW -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_7.svg)

<!-- TBW -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_8.svg)

<!-- TBW -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_9.svg)

<!-- TBW -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_10.svg)

<!-- TBW -->

---
<!-- Erik -->

## By the way: not just artifacts

- Source changes
- Builds
- Test runs
- Failures
- Compositions (multiple artifacts)
- Announcements
- Many many more...

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_X.svg)

<!-- TBW -->



---
<!-- Andrea -->

<!--
_class:
 - lead
 - invert
-->

# Observability & Metrics

<!-- Erik -->

---
<!-- Andrea -->

# TBW Use Case 2 Details

TBW

---
<!-- Andrea -->

<!--
class:
 - lead
 - invert
-->

<!-- Andrea -->

# What can CDEvents do for you?

---
<!-- Andrea -->

# TBW Do for you 1

---
<!-- Andrea -->

# TBW Do for you 2

---
<!-- Andrea -->

# TBW Do for you 3

---
<!-- Andrea -->

# The way forward

---
<!-- Andrea -->

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
<!-- Erik -->

<!--
class:
 - lead
 - invert
-->

<!-- Erik -->

# Key Takeaways

---
<!-- Erik -->

# Spec for interoperability and observability

---
<!-- Erik -->

# TBW Key Takeaway 2

---
<!-- Erik -->

# TBW Key Takeaway 3

---
<!-- Erik -->

# Thank You!

---
<!-- Erik -->

# Questions?

<br />
<br />
<br />

### _Join the cause!_
### _CDEvents <https://github.com/cdevents>_