# scythe : a Harvest API cached implementation in go lang

## Overview [![GoDoc](https://godoc.org/github.com/vanor89/scythe?status.svg)](https://godoc.org/github.com/vanor89/scythe)

This project is meant to be used to save time when logging time repetitively to the same set of projects.

It is called Scythe because a scythe is a tool that is used to harvest quickly. The project will continue to use farming concepts for its implementation
You must configure your account using the `init` command, it will prompt for your personal harvest ID information from the harvest developer site (`https://id.getharvest.com/developers`).

For the time being there are 3 concepts, `ammend`, `sow`, `reap`;

Ammend is what is done to prepare the field for sowing, for our purposes, it fetches the tasks assignments of the currently configured user and overrides the current task assignments in cache if any.

Sow is what is done to the field once to get yields, here is where you start working so for our purposes it is where we will log time.

Reap is what is done to the crop once it has ripened, for our purposes, submitting our time for review and approval.

## Install

```
go get github.com/vanor89/scythe
```

## Contributing

Julio Arias - PR Reviews
`https://github.com/jarias`

## Author

Carlos Villalta
`https://github.com/vanor89`

## License

Copyright 2020 Carlos Villalta

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
