package v1alpha1_test

import (
  "context"
  "github.com/google/go-cmp/cmp"
  "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "testing"
  "time"
)

func TestTaskSpec_SetDefaults(t *testing.T) {
  cases := []struct{
    desc string
    ts *v1alpha1.TaskSpec
    want *v1alpha1.TaskSpec
  }{
    {
      desc: "timeout is nil",
      ts: &v1alpha1.TaskSpec{
        Steps:[]v1alpha1.Step{
          {
            Timeout: nil,
          },
        },
      },
      want:&v1alpha1.TaskSpec{
        Steps:[]v1alpha1.Step{
          {
            Timeout: &metav1.Duration{Duration: 60 * time.Minute},
          },
        },
      },
    },
    {
      desc: "timeout is not nil",
      ts: &v1alpha1.TaskSpec{
        Steps:[]v1alpha1.Step{
          {
            Timeout: &metav1.Duration{Duration: 500 * time.Millisecond},
          },
        },
      },
      want:&v1alpha1.TaskSpec{
        Steps:[]v1alpha1.Step{
          {
            Timeout: &metav1.Duration{Duration: 500 * time.Millisecond},
          },
        },
      },
    },
  }

  for _, tc := range cases {
    t.Run(tc.desc, func(t *testing.T){
      ctx := context.Background()
      tc.ts.SetDefaults(ctx)

      if diff := cmp.Diff(tc.want, tc.ts); diff != "" {
        t.Errorf("Mismatch of TaskSpec: %s", diff)
      }
    })
  }
}