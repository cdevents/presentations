---
marp: true
theme: gaia
backgroundColor: #112
class: invert
_footer: 'FOSDEM CI/CD Devroom | 06.02.2022'
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
Back in November 2019 I was asked by a colleague to help prepare a training session on our continuous delivery setup. *This was for the autonomous drive development at Volvo Cars, so quite a complex project.*

By then we had been working on this continuous delivery setup for about a year and a half and a simplified overview of what we had **built** was needed. After an hour in front of this massive whiteboard drawing boxes for everything we had 72 components on the board, quite a lot but not really surprising since it was quite complex.

Some of these components were configuration files, some were existing solutions like Jenkins, Gerrit and Nexus, but about **25** of the components we identified were basically just **glue code**... Just making things work together... Like moving information from one component to an other, or starting an action in one .

I remember thinking "this work is probably repeated in other projects and organizations too, **there has to be a better way**. And as it turns out, **there is**!
-->

---
<!-- Andrea + Erik -->

<!--
_footer: FOSDEM CI/CD Devroom | 06.02.2022
-->
# CDEvents

### Interoperability in the Continuous Delivery space

<br />
<br />

##### Andrea Frittoli & Erik Sternerson

<!-- Thank you Erik, I loved that story! Indeed we believe there is a better way, and that's what we're going to present today: we will talk about CDevents, a project to provide a protocol and vocabulary for events and signalling in the Continuous Delivery area.
Before we dive in, let's briefly introduce ourselves.

Hi, my name is Andrea Frittoli, I'm a software engineer at IBM. I serve as co-chair for the Events special interest group at the at the continuos delivery foundation and I'm one of the maintainer of the CDEvents specification. I'm also a maintainer of Tekton.

And my name is Erik Sternerson, I work with continuous delivery at doWhile and focus mainly on high-complexity products such as autonomous drive vehicles and mobile network backbones. I am a member of the same Events group that Andrea co-chairs, and I am also a maintainer of the CDEvents specification.
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

<!-- Thanks Andrea! So, first up, we'll look into Continuous Delivery Events as a concept.-->

---
<!-- Erik -->
# A primer on Continuous Delivery

<!-- Just to make sure everyone is on the same page, let's quickly summarize what Continuous Delivery is all about. -->

* ### Goal: Ready-for-release
<!-- The same way that Continuous Integration aims to get your code integrated to the main branch of your repository, Continuous Delivery aims to get your project in a state where it can be released -->
* ### Scope: Idea → Build → Test → Measure → Idea
<!-- In its widest form, Continuous Delivery goes beyond just build and deployment. It can support you all the way from breaking down your initial idea through build and verifiction to getting metrics from your released project, and even using those same metrics as the basis for your next idea. -->
* ### Too much for any one system to solve
<!-- With the wide range of activities needed to realize this scope, it is simple not feasible to have a single "one-size-fits-all" Continuous Delivery system that solves everything. Instead, we need to rely on integrations. -->

---
<!-- Erik -->
![bg contain](images/Continuous-Delivery-Landscape.png)

<!-- Here we see just a few of the open-source projects in the continuous delivery area. Many of these projects will at some point need to be used together for some specific purpose, and this is where the integrations I just mentioned come in. -->

---
<!-- Erik -->
# Continuous Delivery Events

<!-- So, now that we have a shared understanding of continuous delivery, let's move forward to continous delivery events.
There are two fundamental ideas behind this concept. -->

* ### When something valuable happens, __announce__ it
<!-- The first fundamental idea is: When something of value happens, the system that knows about this value announces it by sending an event.

Value is produced in Continuous Delivery all the time. And we are not only talking built binaries and artifacts here. -->
* *Change, Build, Test, Deployment, New Env., Failure, Remediation, etc. etc.*
<!-- It could be everything from code changes, builds. test runs etc. etc. -->
* ### When something interesting is seen, __react__ to it
<!-- The second part of it is that every system that can listen to the events that are sent can chose to react to them.. -->
* *Visualize, Store, Announce, Run, Retry, etc. etc.*
<!-- These reactions can fall into many categories, like storing and visualizing to starting some other activity. -->

---
<!-- Erik -->
# Not a brand new idea

<!-- Now, using events in Continuous Delivery is not something we invented in our group. -->

- ## Tekton
<!-- We have for instance Tekton, which produces CloudEvents and can react to events using triggers. -->

- ## Eiffel Events
<!-- We have Eiffel Events, providing a protocol for many kinds of events related to the production and verification of complex products. -->

- ## Keptn
<!-- And we have Keptn, which uses events to connect to and control services like Helm, Artillery or JMeter. -->

- ## Many more...
<!-- And there are plenty of other systems that send or receive something like events. 

The problem is that the events sent or received by one system is not often understood by others.-->

---
<!-- Erik -->
<!--
class:
 - lead
 - invert
-->

# "Let's work together!"

<!-- So, there was an idea: Let's collaborate on Continuous Delivery Events! -->

---
<!-- Andrea -->
<!--
backgroundColor: #FFF
-->

![bg auto](3rdPartyLogos/cdevents-draft.png)

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

- ## A vocabulary
- ## A specification
- ## A project
- ## In progress :)
- ## [github.com/cdevents](https://github.com/cdevents)

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

- ## Core
- ## Source Code Version Control
- ## Continuous Integration Pipelines
- ## Continuous Deployment Pipelines
- ## ...more to come

<!-- Most CD platform define their abstractions, data model and nomenclature. The interoperability SIG has already been collecting these from various platforms. Many names are shared across platforms, but sometimes the same name bears different meanings in different projects. To achieve interoperability through events we needed to define a nomenclature with shared semantics across platforms. To achieve that we first created four buckets to group the different events. The list is not exhaustive, as we'd like eventually to cover more aspects of the software lifecycle, like for instance monitoring and incident management. Within each bucket, we defined a few abstractions. For instance "Core" defines "Task Runs" and "Pipeline Runs". Continuous integration defines "Build", "Test Case", "Test Suite" and "Artifact". -->

---
<!-- Andrea -->

# The CDEvents specification

- ## Data Model
    - ### Events and Fields
- ## Protocol Bindings
    - ### CloudEvents
- ## Primer

<!-- With the vocabulary defined, the next step has been to define which events are available for each abstraction. For instance a PipelineRun can be queued, started and finished. An artifact and by packaged and published. Each event has mandatory and optional fields. We're also evaluating extension mechanisms, to allow two kind of extension: documented third party fields as well as free-style content.

Since there are plenty of established messaging protocols already available, we didn't want to reinvent the wheel. Rather instead we have been working at identifying which fields are required for the continuous delivery use cases, and then we defined protocol bindings which define how CDEvents map onto a specific protocol.
CloudEvents - a CNCF project - are an established format in the Cloud Native space, and that's the first binding we defined.

Finally, as part of the specification we started working on a cdevents primer document, which documents the protocol architecture, design decisions and reference use cases. -->

---
<!-- Andrea -->

<!--
class:
 - lead
 - invert
-->

# CDEvents use cases

<!-- Use cases are key to development of cdevents. When we define events and their fields, we ask the question "what is a minimal set of information we need to satisfy the use cases?"
There are two root use cases, which we identified until now.
The first one is is interoperability, in the sense on making it possible for a CD platform to consume events produced by another one, without the need of dedicated glue code.
The second one is observability and metrics. In a way this is also about interoperability, and it refers to the ability of collecting events from different CD platforms, being able to correlate them and process them consistently to build and end-to-end view of the overall CD workflow an its properties, like duration.
-->

---
<!-- Erik -->

# Interoperability

<!-- The first use case we want to talk about today is making things work together.  -->

---
<!-- Erik -->

<!--
class:
 - invert
-->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_1.svg)

<!-- I would like to take you with me back to my world for a while, where I am on a quest to build a car. A modern car has hundreds of software modules in it, and actually getting from an individual software module to a working car is not trivial. For this use case, we will focus in on builds. -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_2.svg)

<!-- In this organization we want to let teams choose their own optimal CI/CD set-up, and many of the software development teams in this organization prefer to use Zuul, for its dependency handling and scalability. -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_3.svg)

<!-- Some teams prefer GoCD for the way it visualizes deep dependency chains. -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_4.svg)

<!-- Some teams prefer GitLab to have both the source code and the project stuff in the same platform.-->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_5.svg)

<!-- Some teams prefer Jenkins for the ease-of-use and the wide variety of plugins. -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_6.svg)

<!-- To further complicate things, we don't build all the software in-house, we also have suppliers. -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_7.svg)

<!-- But all these services and suppliers produce built software modules for us. -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_8.svg)

<!-- We use these built software moduloes for various purposes, for instance running them in the audonomous drive simulator plaftorm Carla, so we can test them before we "load" them into the car itself. -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_9.svg)

<!-- So now we start looking at the Zuul docs to see what the output of a build is. An "Artifact". Right. For GoCD? "Artifact". GitLab and Jenkins? Also "Artifacts". Then we are home free right?   -->

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_10.svg)

<!-- Well, no. Because a Zuul artifact is a bit different from a GoCD artifact, which is a bit different from a GitLab artifact, and so on an so forth. So we will likely need to write a bit of translation or glue code to be able to actually understand and receive all these built software modules. -->

---
<!-- Erik -->

## By the way: not just artifacts

<!-- And this really doesn't apply only to artifacts. -->

- Source changes
- Build activities
- Test runs
- Failures
- Compositions (multiple artifacts)
- Announcements
- Many many more...

<!-- So, what would we want instead? -->

---
<!-- Erik -->

![bg contain](images/glue_code.jpg)

<!-- And we are basically back here again, writing more and more glue code to fit things together. -->

---
<!-- Erik -->


<!--
_class:
 - lead
 - invert
-->

# Solution: CDEvents

---
<!-- Erik -->

![bg contain](images/CDEvents-FOSDEM2022-Interoperability_X.svg)

<!-- We would like all these build systems, and our suppliers, to have the option to announce new artifacts in a standardized format. CDEvents will provide this format, but will need support from the community to start applying it to relevant projects. -->

---
<!-- Andrea -->

<!--
_class:
 - lead
 - invert
-->

# Observability & Metrics

<!-- Thanks Erik, that was a really good use case. The next one is observability and metrics -->
---
<!-- Andrea -->

![bg contain](images/cdevents_use_case_observability_1.svg)

<!-- Let's consider a the following CD setup, where code is written and maintained in GitHub. When changes are made, they go through different tests, maintained by different teams, which use different technologies. Some tests are running in GitHub directly as GitHub actions. Some others are executed in Jenkins and other as Tekton pipelines. Releases are managed through Tekton as well, while deployments are managed with Argo. Keptn is used to manage remediations strategies on production clusters.
-->

---
<!-- Andrea -->

![bg contain](images/cdevents_use_case_observability_2.svg)

<!-- How may we visualize an e2e workflow?
All these systems supports events in some format, so we could start collecting events. Since but they are all different though, our events collector would need to support multiple ways of collecting events. Both Tekon and Keptn use cloud events, which makes things a bit easier, but even for those two, there are no shared semantics.-->

---
<!-- Andrea -->

![bg contain](images/cdevents_use_case_observability_3.svg)

<!-- If all platforms supported the same format of events, like CloudEvents, our event collector could become one of the existing events broker, like Knative Eventing. Going back to the e2e workflow though, if we wanted to visualize the flow of a change from when it's written, through test, release, deploy and possibly rollback, we'd need to have enough information in the events to be able to correlated with each other. This is one if the issues that we want to address with CDEvents.
Another question that we may want to be able to answer is how effective is the DevOps setup of my organisation? To answer that you would need to define metrics, collect the relevant data and visualise it. CDEvents may help collecting the data from heterogeneous sources, store it and process it in a consistent way.
How do we know what kind of data we need though? In the cdevents project we started considering the metrics defined by DORA (DevOps research and assesment) as a reference to validate against.
-->
---
<!-- Andrea -->

<!--
class:
 - lead
 - invert
-->

<!-- Andrea -->

# What can CDEvents do for you?

<!-- To summarise, what can CDEvents do for you?
Events let you build a decoupled, scalable and reliable architecture. 
As you create your build your CD workflows, CDEvents will let you: -->
---
<!-- Andrea -->

<!--
class:
 - invert
-->

![bg right:60%](images/glue_code.jpg)

## Reduce Glue Code
---
<!-- Andrea -->

![bg left:66%](images/mike-benna-X-NAMq6uP3Q-unsplash.jpg)
## End to End CD Workflow Visibility
---
<!-- Andrea -->

![bg right:60%](images/william-warby-WahfNoqbYnM-unsplash.jpg)
## Collect data to measure DevOps Performance
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

<!-- We've reached the end of our talk, and we wanted to summarize a bit before we move on to the Q&A section. -->

---
<!-- Erik -->

# CDEvents
## A spec for interoperability and observability

<!-- The CDEvents project, with contributors from all over the CI/CD community, is building a specification for events in CI/CD that can be used to enable both interoperability and observability for Continuous Delivery systems. -->

---
<!-- Erik -->

# Integrations needed

<!-- Of course a spec that no one uses is not benefiting anyone, so we will need both SDKs for sending events and integrations into the projects that could send and listen to these events. Could be a good starter-task for new developers to your project, or maybe something for a GSoC student? -->

---
<!-- Erik -->

# We want to hear from YOU!

[tinyurl.com/talktocdevents](https://tinyurl.com/talktocdevents)
<!-- Fundamentally the purpose of CDEvents is to become something that helps YOU. So, something we really need right now is to learn more about the pain points and needs that YOU have in your day-to-day, and what you see as upcoming challenges. Please get in touch! -->

---
<!-- Erik -->

# Thank You!

---
<!-- Erik -->
<!--
_footer: 'FOSDEM CI/CD Devroom | 06.02.2022'
-->

# Questions?

<br />
<br />
<br />

### _Join the cause!_
### _CDEvents <https://github.com/cdevents>_
### [tinyurl.com/talktocdevents](https://tinyurl.com/talktocdevents)
