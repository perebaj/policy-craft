import { DiamondSvg } from 'assets/Diamond'
import * as React from 'react'
import { Controller, SubmitHandler, useForm } from 'react-hook-form'
import { Handle, NodeProps, Position } from 'reactflow'

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
import { Input } from '@/src/components/ui/input'
import { Label } from '@/src/components/ui/label'
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

// ConditionalNode is used to render the conditional node
export function ConditionalNode({ data }: NodeProps<ConditionalNodeData>) {
  const [open, setOpen] = React.useState(false)

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
          <ConditionalForm />
        </DialogContent>
      </Dialog>
    </NodeWrapper>
  )
}

// ConditionalFormInputs is used to define the user input for the conditional node
type ConditionalFormInputs = {
  name: string
  criteria: string
  compvalue: number
}

// ConditionalForm is used to capture the user input for the conditional node
function ConditionalForm({ className }: React.ComponentProps<'form'>) {
  const { register, handleSubmit, control } = useForm<ConditionalFormInputs>()
  const onSubmit: SubmitHandler<ConditionalFormInputs> = (data) =>
    console.log(data)

  return (
    <form
      className={cn('grid items-start gap-4', className)}
      onSubmit={handleSubmit(onSubmit)}
    >
      <div className="grid w-full gap-2">
        <Label htmlFor="name">Component Name</Label>
        <Input type="text" id="name" {...register('name')} />
      </div>
      <div className="flex w-full items-center justify-center gap-4">
        <div className="w-full">
          <Label htmlFor="criteria">Comparation Criteria</Label>
          {/* Controller is used to capture the user input in the form for this
          specific input field (Select) */}
          <Controller
            render={({ field }) => (
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
            )}
            name="criteria"
            control={control}
          />
        </div>
        <div className="w-full">
          <Label htmlFor="compvalue">Comparation Value</Label>
          <Input type="number" id="compvalue" {...register('compvalue')} />
        </div>
      </div>
      <Button type="submit" className="bg-cyan-500 hover:bg-cyan-400">
        Save changes
      </Button>
    </form>
  )
}
