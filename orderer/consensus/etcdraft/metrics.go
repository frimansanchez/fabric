/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package etcdraft

import "github.com/hyperledger/fabric/common/metrics"

var (
	clusterSizeOpts = metrics.GaugeOpts{
		Namespace:    "consensus",
		Subsystem:    "etcdraft",
		Name:         "cluster_size",
		Help:         "Number of nodes in this channel.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
	isLeaderOpts = metrics.GaugeOpts{
		Namespace:    "consensus",
		Subsystem:    "etcdraft",
		Name:         "is_leader",
		Help:         "The leadership status of the current node: 1 if it is the leader else 0.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
	committedBlockNumberOpts = metrics.GaugeOpts{
		Namespace:    "consensus",
		Subsystem:    "etcdraft",
		Name:         "committed_block_number",
		Help:         "The block number of the latest block committed.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
	snapshotBlockNumberOpts = metrics.GaugeOpts{
		Namespace:    "consensus",
		Subsystem:    "etcdraft",
		Name:         "snapshot_block_number",
		Help:         "The block number of the latest snapshot.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
	leaderChangesOpts = metrics.CounterOpts{
		Namespace:    "consensus",
		Subsystem:    "etcdraft",
		Name:         "leader_changes",
		Help:         "The number of leader changes.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
	proposalFailuresOpts = metrics.CounterOpts{
		Namespace:    "consensus",
		Subsystem:    "etcdraft",
		Name:         "proposal_failures",
		Help:         "The number of proposal failures.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
)

type Metrics struct {
	ClusterSize          metrics.Gauge
	IsLeader             metrics.Gauge
	CommittedBlockNumber metrics.Gauge
	SnapshotBlockNumber  metrics.Gauge
	LeaderChanges        metrics.Counter
	ProposalFailures     metrics.Counter
}

func NewMetrics(p metrics.Provider) *Metrics {
	return &Metrics{
		ClusterSize:          p.NewGauge(clusterSizeOpts),
		IsLeader:             p.NewGauge(isLeaderOpts),
		CommittedBlockNumber: p.NewGauge(committedBlockNumberOpts),
		SnapshotBlockNumber:  p.NewGauge(snapshotBlockNumberOpts),
		LeaderChanges:        p.NewCounter(leaderChangesOpts),
		ProposalFailures:     p.NewCounter(proposalFailuresOpts),
	}
}