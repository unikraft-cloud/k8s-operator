// SPDX-License-Identifier: UNLICENSED
//
// Copyright (c) 2025, Unikraft GmbH.  All rights reserved.
//
// This software and related documentation ("Unikraft Software") are protected
// under relevant copyright laws.  The information contained herein is
// confidential and proprietary to Unikraft GmbH ("Unikraft") and/or its
// licensors.  Without the prior written permission of Unikraft GmbH and/or its
// licensors, any reproduction, modification, use or disclosure of Unikraft
// Software, and information contained herein, in whole or in part, shall be
// strictly prohibited.
//
// BY OPENING THIS FILE, RECEIVER HEREBY UNEQUIVOCALLY ACKNOWLEDGES AND AGREES
// THAT THE SOFTWARE/FIRMWARE AND ITS DOCUMENTATIONS ("UNIKRAFT SOFTWARE")
// RECEIVED FROM UNIKRAFT AND/OR ITS REPRESENTATIVES ARE PROVIDED TO RECEIVER ON
// AN "AS-IS" BASIS ONLY.  UNIKRAFT EXPRESSLY DISCLAIMS ANY AND ALL WARRANTIES,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE IMPLIED WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE OR NONINFRINGEMENT. NEITHER
// DOES UNIKRAFT PROVIDE ANY WARRANTY WHATSOEVER WITH RESPECT TO THE SOFTWARE OF
// ANY THIRD PARTY WHICH MAY BE USED BY, INCORPORATED IN, OR SUPPLIED WITH THE
// UNIKRAFT SOFTWARE, AND RECEIVER AGREES TO LOOK ONLY TO SUCH THIRD PARTY FOR
// ANY WARRANTY CLAIM RELATING THERETO.  RECEIVER EXPRESSLY ACKNOWLEDGES THAT IT
// IS RECEIVER'S SOLE RESPONSIBILITY TO OBTAIN FROM ANY THIRD PARTY ALL PROPER
// LICENSES CONTAINED IN UNIKRAFT SOFTWARE.  UNIKRAFT SHALL ALSO NOT BE
// RESPONSIBLE FOR ANY UNIKRAFT SOFTWARE RELEASES MADE TO RECEIVER'S
// SPECIFICATION OR TO CONFORM TO A PARTICULAR STANDARD OR OPEN FORUM.
// RECEIVER'S SOLE AND EXCLUSIVE REMEDY AND UNIKRAFT'S ENTIRE AND CUMULATIVE
// LIABILITY WITH RESPECT TO THE UNIKRAFT SOFTWARE RELEASED HEREUNDER WILL BE,
// AT UNIKRAFT'S OPTION, TO REVISE OR REPLACE THE UNIKRAFT SOFTWARE AT ISSUE, OR
// REFUND ANY SOFTWARE LICENSE FEES OR SERVICE CHARGE PAID BY RECEIVER TO
// UNIKRAFT FOR SUCH UNIKRAFT SOFTWARE AT ISSUE.

package main

import (
	"crypto/tls"

	"github.com/alecthomas/kong"
)

type Config struct {
	UKCToken string `name:"ukc-token" env:"OPERATOR_UKC_TOKEN" help:"Unikraft Cloud access token." required:""`
	UKCMetro string `name:"ukc-metro" env:"OPERATOR_UKC_METRO" help:"Default Unikraft Cloud metro where the instances will be created." required:""`

	MetricsAddr             string              `name:"metrics-addr" env:"METRICS_ADDR" help:"The address the metric endpoint binds to." default:":8080"`
	ProbeAddr               string              `name:"probe-addr" env:"PROBE_ADDR" help:"The address the probe endpoint binds to." default:":8081"`
	EnableLeaderElection    bool                `name:"enable-leader-election" env:"ENABLE_LEADER_ELECTION" help:"Enable leader election for controller manager." default:"false"`
	LeaderElectionID        string              `name:"leader-election-id" env:"LEADER_ELECTION_ID" help:"ID of the leader election lease." default:"operator.unikraft.com"`
	LeaderElectionNamespace string              `name:"leader-election-namespace" env:"LEADER_ELECTION_NAMESPACE" help:"Namespace of the operator lease." default:"ukc-operator"`
	SecureMetrics           bool                `name:"metrics-secure" env:"METRICS_SECURE" help:"If set, the metrics endpoint is served securely via HTTPS. Use --metrics-secure=false to use HTTP instead." default:"true"`
	EnableHTTP2             bool                `name:"enable-http2" env:"ENABLE_HTTP2" help:"If set, HTTP/2 will be enabled for the metrics and webhook servers" default:"false"`
	TLSOpts                 []func(*tls.Config) `name:"" env:"" `
}

func parseConfig() *Config {
	var cfg Config
	kong.Parse(&cfg,
		kong.Name("ukc-k8s-operator"),
		kong.Description("Unikraft Cloud Kubernetes operator."),
	)
	return &cfg
}
