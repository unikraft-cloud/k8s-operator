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

package controller

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	unikraftv1alpha1 "github.com/unikraft-cloud/k8s-operator/api/v1alpha1"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var (
	cfg       *rest.Config
	k8sClient client.Client
	testEnv   *envtest.Environment
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Controller Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "..", "config", "crd", "bases")},
		ErrorIfCRDPathMissing: true,

		// The BinaryAssetsDirectory is only required if you want to run the tests directly
		// without call the makefile target test. If not informed it will look for the
		// default path defined in controller-runtime which is /usr/local/kubebuilder/.
		// Note that you must have the required binaries setup under the bin directory to perform
		// the tests directly. When we run make test it will be setup and used automatically.
		BinaryAssetsDirectory: filepath.Join("..", "..", "bin", "k8s",
			fmt.Sprintf("1.30.0-%s-%s", runtime.GOOS, runtime.GOARCH)),
	}

	var err error
	// cfg is defined in this file globally.
	cfg, err = testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	err = unikraftv1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	// +kubebuilder:scaffold:scheme

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).NotTo(BeNil())
})

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	err := testEnv.Stop()
	Expect(err).NotTo(HaveOccurred())
})
