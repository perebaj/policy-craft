import { zodResolver } from '@hookform/resolvers/zod'
import { DiamondSvg } from 'assets/Diamond'
import * as React from 'react'
import { SubmitHandler, useForm, UseFormReturn } from 'react-hook-form'
import { Handle, NodeProps, Position } from 'reactflow'
import { z } from 'zod'

import { cn } from '@/lib/utils'
import { Button } from '@/src/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/src/components/ui/dialog'
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/src/components/ui/form'
import { Input } from '@/src/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/src/components/ui/select'

import { NodeWrapper } from './NodeWrapper'

export type ConditionalNodeData = {
  label: string
  elseLabel?: string
  width: number
  height: number
}

// formSchema is used to validate the user input using the zod library
const formSchema = z.object({
  name: z.string().min(1),
  criteria: z.string().min(1),
  compValue: z.number(),
})

// ConditionalNode is used to render the conditional node
export function ConditionalNode({ data }: NodeProps<ConditionalNodeData>) {
  const [open, setOpen] = React.useState(false)
  async function onSubmit(value: z.infer<typeof formSchema>) {
    console.log(value)
    setOpen(false)
  }

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: '',
      criteria: '',
      compValue: 0,
    },
  })

  return (
    <NodeWrapper>
      <Dialog open={open} onOpenChange={setOpen}>
        <div
          className="group h-full w-full"
          style={{
            width: data.width,
            height: data.height,
          }}
        >
          <DialogTrigger asChild>
            {/* <Button variant="outline">Edit </Button> */}
            <div className="relative flex h-full w-full items-center justify-center   p-9 text-center text-[12px]">
              <div className="group-hover-focus:cursor-pointer absolute left-0 top-0 z-0 h-full w-full text-cyan-400 hover:cursor-pointer group-hover:text-cyan-400/80 [&>svg]:stroke-cyan-500">
                <DiamondSvg strokeWidth={4} />
              </div>
              <Handle
                type="target"
                id="target"
                className="invisible"
                position={Position.Top}
                isConnectable={false}
              />
              <p className={`z-10 line-clamp-3 cursor-pointer`}>{data.label}</p>
              <Handle
                type="source"
                id="source"
                className="invisible"
                position={Position.Bottom}
                isConnectable={false}
              />
            </div>
          </DialogTrigger>
        </div>
        <DialogContent className="max-h-5xl">
          <DialogHeader>
            <DialogTitle>Edit your conditional component</DialogTitle>
            <DialogDescription>
              Make changes to your conditional component. Click save when you
              are done.
            </DialogDescription>
          </DialogHeader>
          <ConditionalForm form={form} onSubmit={onSubmit} />
        </DialogContent>
      </Dialog>
    </NodeWrapper>
  )
}

// ConditionalForm is used to capture the user input for the conditional node
// @param form - useForm from react-hook-form to capture the user input
// @param onSubmit - function to be called when the form is submitted and close the dialog component that wraps the form
function ConditionalForm({
  form,
  onSubmit,
}: {
  form: UseFormReturn<z.infer<typeof formSchema>>
  onSubmit: SubmitHandler<z.infer<typeof formSchema>>
}) {
  return (
    <Form {...form}>
      <form
        className={cn('grid items-start gap-4')}
        onSubmit={form.handleSubmit(onSubmit)}
      >
        <div className="grid w-full gap-2">
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel htmlFor="name">Component Name</FormLabel>
                <FormControl>
                  <Input type="text" id="name" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>
        <div className="flex w-full items-center justify-center gap-4">
          <div className="w-full">
            <FormField
              control={form.control}
              name="criteria"
              render={({ field }) => (
                <FormItem>
                  <FormLabel htmlFor="criteria">Comparation Criteria</FormLabel>
                  <Select value={field.value} onValueChange={field.onChange}>
                    <SelectTrigger className="w-full">
                      <SelectValue placeholder="Comparation Criteria" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectItem value=">">{`>`}</SelectItem>
                      <SelectItem value=">=">{`>=`}</SelectItem>
                      <SelectItem value="<">{`<`}</SelectItem>
                      <SelectItem value="<=">{`<=`}</SelectItem>
                      <SelectItem value="==">{`==`}</SelectItem>
                    </SelectContent>
                  </Select>
                </FormItem>
              )}
            />
          </div>
          <div className="w-full">
            <FormField
              control={form.control}
              name="compValue"
              render={({ field }) => (
                <FormItem>
                  <FormLabel htmlFor="CompValue">Comparation Value</FormLabel>
                  <FormControl>
                    <Input
                      type="number"
                      id="CompValue"
                      {...field}
                      onChange={(event) => field.onChange(+event.target.value)}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
        </div>
        <Button type="submit" className="bg-cyan-500 hover:bg-cyan-400">
          Save changes
        </Button>
      </form>
    </Form>
  )
}
