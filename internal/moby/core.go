// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package moby

import "github.com/docker/docker/client"

type moby struct {
	client *client.Client
}

func (m *moby) Whoami() string {
	return "docker(moby)"
}
