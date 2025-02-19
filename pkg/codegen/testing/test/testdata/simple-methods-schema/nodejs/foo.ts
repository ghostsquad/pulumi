// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

import * as pulumiRandom from "@pulumi/random";

export class Foo extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'example::Foo';

    /**
     * Returns true if the given object is an instance of Foo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Foo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Foo.__pulumiType;
    }


    /**
     * Create a Foo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FooArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
        } else {
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Foo.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }

    /**
     * A description of bar.
     */
    bar(args: Foo.BarArgs): pulumi.Output<Foo.BarResult> {
        return pulumi.runtime.call("example::Foo/bar", {
            "__self__": this,
            "baz": args.baz,
            "bazPlain": args.bazPlain,
            "bazRequired": args.bazRequired,
            "boolValue": args.boolValue,
            "boolValuePlain": args.boolValuePlain,
            "boolValueRequired": args.boolValueRequired,
            "name": args.name,
            "namePlain": args.namePlain,
            "nameRequired": args.nameRequired,
            "stringValue": args.stringValue,
            "stringValuePlain": args.stringValuePlain,
            "stringValueRequired": args.stringValueRequired,
        }, this);
    }

    baz(): void {
        pulumi.runtime.call("example::Foo/baz", {
            "__self__": this,
        }, this);
    }

    /**
     * Do something with something else
     */
    generateKubeconfig(args: Foo.GenerateKubeconfigArgs): pulumi.Output<Foo.GenerateKubeconfigResult> {
        return pulumi.runtime.call("example::Foo/generateKubeconfig", {
            "__self__": this,
            "boolValue": args.boolValue,
        }, this);
    }
}

/**
 * The set of arguments for constructing a Foo resource.
 */
export interface FooArgs {
}

export namespace Foo {
    /**
     * The set of arguments for the Foo.bar method.
     */
    export interface BarArgs {
        baz?: pulumi.Input<inputs.nested.BazArgs>;
        bazPlain?: inputs.nested.Baz;
        bazRequired: pulumi.Input<inputs.nested.BazArgs>;
        boolValue?: pulumi.Input<boolean>;
        boolValuePlain?: boolean;
        boolValueRequired: pulumi.Input<boolean>;
        name?: pulumi.Input<pulumiRandom.RandomPet>;
        namePlain?: pulumiRandom.RandomPet;
        nameRequired: pulumi.Input<pulumiRandom.RandomPet>;
        stringValue?: pulumi.Input<string>;
        stringValuePlain?: string;
        stringValueRequired: pulumi.Input<string>;
    }

    /**
     * The results of the Foo.bar method.
     */
    export interface BarResult {
        readonly someValue: string;
    }

    /**
     * The set of arguments for the Foo.generateKubeconfig method.
     */
    export interface GenerateKubeconfigArgs {
        boolValue: boolean;
    }

    /**
     * The results of the Foo.generateKubeconfig method.
     */
    export interface GenerateKubeconfigResult {
        readonly kubeconfig: string;
    }

}
