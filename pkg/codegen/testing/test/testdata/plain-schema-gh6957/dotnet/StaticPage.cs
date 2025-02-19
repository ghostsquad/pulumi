// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Xyz
{
    [XyzResourceType("xyz:index:StaticPage")]
    public partial class StaticPage : Pulumi.ComponentResource
    {
        /// <summary>
        /// The bucket resource.
        /// </summary>
        [Output("bucket")]
        public Output<Pulumi.Aws.S3.Bucket> Bucket { get; private set; } = null!;

        /// <summary>
        /// The website URL.
        /// </summary>
        [Output("websiteUrl")]
        public Output<string> WebsiteUrl { get; private set; } = null!;


        /// <summary>
        /// Create a StaticPage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StaticPage(string name, StaticPageArgs args, ComponentResourceOptions? options = null)
            : base("xyz:index:StaticPage", name, args ?? new StaticPageArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class StaticPageArgs : Pulumi.ResourceArgs
    {
        [Input("foo")]
        public Inputs.Foo? Foo { get; set; }

        /// <summary>
        /// The HTML content for index.html.
        /// </summary>
        [Input("indexContent", required: true)]
        public Input<string> IndexContent { get; set; } = null!;

        public StaticPageArgs()
        {
        }
    }
}
